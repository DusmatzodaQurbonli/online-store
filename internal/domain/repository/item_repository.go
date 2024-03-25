package repository

import (
	"github.com/DusmatzodaQurbonli/online-store/internal/infrastructure/db"
)

type ItemRepository struct {
	db *db.Database
}

func NewItemRepository(db *db.Database) ItemRepository {
	return &ItemRepository{db: db}
}
