package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/DusmatzodaQurbonli/online-store/internal/services"
	"github.com/DusmatzodaQurbonli/online-store/internal/types"
)

type Handler struct {
	Service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetOrders(orderIDs []string) ([]types.Item, error) {
	return h.Service.GetOrdersByID(context.Background(), orderIDs)
}

func (h *Handler) PrintItems(orderIDs []string, items []types.Item) {
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
