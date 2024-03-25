package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"

	"github.com/DusmatzodaQurbonli/online-store/internal/db"
	"github.com/DusmatzodaQurbonli/online-store/internal/handler"
	"github.com/DusmatzodaQurbonli/online-store/internal/services"
	"github.com/DusmatzodaQurbonli/online-store/internal/types"

)

const (
	host     = "localhost"
	port     = 5432
	user     = "app"
	password = "pass"
	dbname   = "db"
)

func main() {
	const pathConfig = `config/config.json`
	config, err := getConfig(pathConfig)
	if err != nil {
		fmt.Println("Error getting config:", err)
		return
	}
	newDB, err := db.NewDB(config)
	if err != nil {
		fmt.Println("Error getting newdb:", err)
		return
	}
	defer newDB.Pool.Close()
	service := services.NewService(newDB)
	handler := handler.NewHandler(service)

	orderIDs := strings.Split(os.Args[1], ",")
	items, err := handler.GetOrders(orderIDs)
	if err != nil {
		fmt.Println("Error getting items:", err)
		return
	}

	handler.PrintItems(orderIDs, items)
}

func getConfig(path string) (*types.Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var config types.Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &config, nil
}
