package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/lib/pq"
)

type Item struct {
	Shelf     string
	Name      string
	Order     int
	Amount    int
	Extra     string
	ProductID int
}

const (
	host     = "localhost"
	port     = 5432
	user     = "app"
	password = "pass"
	dbname   = "db"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	orderIDs := strings.Split(os.Args[1], ",")
	rows, err := db.Query(`
		SELECT s.name, p.name, p.id, po.order_id, po.quantity,
		COALESCE((SELECT STRING_AGG(s2.name, ',') FROM ProductsShelves ps2 JOIN Shelves s2 ON s2.id = ps2.shelf_id WHERE ps2.product_id = p.id AND ps2.is_main = false), '') as extra
		FROM ProductsOrders po
		JOIN Products p ON po.product_id = p.id
		JOIN ProductsShelves ps ON ps.product_id = p.id
		JOIN Shelves s ON s.id = ps.shelf_id
		WHERE po.order_id = ANY($1) AND ps.is_main = true
		ORDER BY s.name, po.order_id ASC
	`, pq.Array(orderIDs))
	if err != nil {
		fmt.Println("Error executing the query:", err)
		return
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.Shelf, &item.Name, &item.ProductID, &item.Order, &item.Amount, &item.Extra); err != nil {
			fmt.Println("Error reading the row:", err)
			return
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error reading the rows:", err)
		return
	}

	fmt.Println("=+=+=+=\nСтраница сборки заказов", strings.Join(orderIDs, ","))
	currentShelf := ""
	for _, item := range items {
		if item.Shelf != currentShelf {
			fmt.Println("===Стеллаж", item.Shelf)
			currentShelf = item.Shelf
		}

		fmt.Printf("%s (id=%d)\nзаказ %d, %d шт", item.Name, item.ProductID, item.Order, item.Amount)

		if item.Extra != "" {
			fmt.Printf(", \nдоп стеллаж: %s\n", item.Extra)
		} else {
			fmt.Println()
		}
		fmt.Println()
	}
}
