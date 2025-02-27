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

// func OrderAddProduct(ctx context.Context, repo repository.Querier, product string) ([]*repository.OrderUserRow, error) {
// 	if err := repo.ProductUpdate(ctx, product); err != nil {
// 		return nil, err
// 	}
// 	product, err := ProductByName(ctx, repo, newProduct.Name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return product, nil
// 	return nil, nil
// }
