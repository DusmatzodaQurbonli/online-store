package db

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	errors "github.com/pkg/errors"

	"github.com/DusmatzodaQurbonli/online-store/internal/types"
)

type DB struct {
	Pool *pgxpool.Pool
}

func NewDB(config *types.Config) (*DB, error) {
	dsn := "postgres://" + config.UserName + ":" + config.Password + "@" + config.Host + ":" + config.Port + "/" + config.Database
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		log.Println(err)
		return nil, errors.Wrap(err, "pgx fail connect")
	}
	return &DB{Pool: pool}, nil
}

func (db *DB) GetOrdersByID(ctx context.Context, orderIDs []string) ([]types.Item, error) {
	items := []types.Item{}

	// Запрос 1: Получаем основные данные о продуктах
	rows, err := db.Pool.Query(ctx, `
	SELECT po.product_id, po.order_id, po.quantity
	FROM ProductsOrders po
	WHERE po.order_id = ANY($1)
`, pq.Array(orderIDs))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item types.Item
		if err := rows.Scan(&item.ProductID, &item.Order, &item.Amount); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	// Запрос 2: Получаем данные о продуктах
	for i, item := range items {
		row := db.Pool.QueryRow(ctx, `
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
		row := db.Pool.QueryRow(ctx, `
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
		rows, err := db.Pool.Query(ctx, `
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
