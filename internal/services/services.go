package services

import (
	"context"

	"github.com/DusmatzodaQurbonli/online-store/internal/db"
	"github.com/DusmatzodaQurbonli/online-store/internal/types"

)

type Service struct {
	db *db.DB
}

func NewService(db *db.DB) *Service {
	return &Service{db: db}
}

func (s *Service) GetOrdersByID(ctx context.Context, orderIDs []string) ([]types.Item, error) {
	return s.db.GetOrdersByID(ctx, orderIDs)
}
