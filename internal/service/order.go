package service

import (
	"context"

	"github.com/henriquedessen/fullcycle-clean_architecture/internal/model"
	"github.com/henriquedessen/fullcycle-clean_architecture/internal/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) ListOrders(ctx context.Context) ([]model.Order, error) {
	return s.repo.ListOrders(ctx)
}
