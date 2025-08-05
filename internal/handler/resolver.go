package handler

import (
	"context"

	"github.com/henriquedessen/fullcycle-clean_architecture/internal/model"
)

type QueryResolver interface {
	ListOrders(ctx context.Context) ([]*model.Order, error)
}
