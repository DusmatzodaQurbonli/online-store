package db

import (
	"database/sql"
	"strings"

	"github.com/lib/pq"

	"github.com/DusmatzodaQurbonli/online-store/internal/domain/entity"

)

type Database struct {
	*sql.DB
}

func NewDatabase(connStr string) (*Database, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

func (p *Database) GetItems(orderIDs []string) ([]entity.Item, error) {
	items := []entity.Item{}

	// Запрос 1: Получаем основные данные о продуктах
	rows, err := p.DB.Query(`
		SELECT po.product_id, po.order_id, po.quantity
		FROM ProductsOrders po
		WHERE po.order_id = ANY($1)
	`, pq.Array(orderIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item entity.Item
		if err := rows.Scan(&item.ProductID, &item.Order, &item.Amount); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	// Запрос 2: Получаем данные о продуктах
	for i, item := range items {
		row := p.DB.QueryRow(`
			SELECT p.name
			FROM Products p
			WHERE p.id = $1
		`, item.ProductID)

		var name string
		if err := row.Scan(&name); err != nil {
			return nil, err
		}
		items[i].Name = name
	}

	// Запрос 3: Получаем данные о полках
	for i, item := range items {
		row := p.DB.QueryRow(`
			SELECT s.name
			FROM Shelves s
			WHERE s.id = (
				SELECT ps.shelf_id
				FROM ProductsShelves ps
				WHERE ps.product_id = $1 AND ps.is_main = true
			)
		`, item.ProductID)

		var shelf string
		if err := row.Scan(&shelf); err != nil {
			return nil, err
		}
		items[i].Shelf = shelf
	}

	// Запрос 4: Получаем дополнительные данные о полках
	for i, item := range items {
		rows, err := p.DB.Query(`
			SELECT s.name
			FROM Shelves s
			WHERE s.id IN (
				SELECT ps.shelf_id
				FROM ProductsShelves ps
				WHERE ps.product_id = $1 AND ps.is_main = false
			)
		`, item.ProductID)

		if err != nil {
			return nil, err
		}

		var extras []string
		for rows.Next() {
			var extra string
			if err := rows.Scan(&extra); err != nil {
				return nil, err
			}
			extras = append(extras, extra)
		}
		items[i].Extra = strings.Join(extras, ",")
	}

	return items, nil
}
