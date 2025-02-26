package service

import (
	"context"

	"github.com/sar0868/otus_go_basic_hw/hw15_go_sql/internal/repository"
)

func Orders(ctx context.Context, repo repository.Querier) ([]*repository.OrdersRow, error) {
	orders, err := repo.Orders(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
