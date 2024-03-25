package handler

import (
	"fmt"
	"strings"
)

type ItemHandler struct {
	usecase usecase.ItemUsecase
}

func NewItemHandler(usecase usecase.ItemUsecase) *ItemHandler {
	return &ItemHandler{usecase: usecase}
}

func (h *ItemHandler) GetItems(orderIDs []string) ([]entity.Item, error) {
	return h.usecase.GetItems(orderIDs)
}

func (h *ItemHandler) PrintItems(orderIDs []string, items []entity.Item) {
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
	}
}
