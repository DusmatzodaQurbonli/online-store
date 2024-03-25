package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/DusmatzodaQurbonli/online-store/internal/app/handler"
	"github.com/DusmatzodaQurbonli/online-store/internal/app/usecase"
	"github.com/DusmatzodaQurbonli/online-store/internal/domain/repository"
	"github.com/DusmatzodaQurbonli/online-store/internal/infrastructure/db"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "app"
	password = "pass"
	dbname   = "db"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := db.NewDatabase(connStr)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	itemRepo := repository.NewItemRepository(db)
	itemUsecase := usecase.NewItemUsecase(itemRepo)
	itemHandler := handler.NewItemHandler(itemUsecase)

	orderIDs := strings.Split(os.Args[1], ",")
	items, err := itemHandler.GetItems(orderIDs)
	if err != nil {
		fmt.Println("Error getting items:", err)
		return
	}

	itemHandler.PrintItems(items)
}
