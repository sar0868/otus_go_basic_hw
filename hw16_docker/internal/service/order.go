package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sar0868/otus_go_basic_hw/hw16_docker/internal/repository"
)

func Orders(ctx context.Context, repo repository.Querier) ([]*repository.OrdersRow, error) {
	orders, err := repo.Orders(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

type OrdProdUpParams struct {
	OrderID  int
	Product  string
	Quantity float64
}

func OrderProductUpdate(ctx context.Context, repo repository.Querier, params repository.OrderProductUpdateParams) (*repository.ShopOrder, error) { //nolint: lll
	err := repo.OrderProductUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	fmt.Println(params)
	order, errResult := OrderRecount(ctx, repo, int(params.OrderID))
	if errResult != nil {
		return nil, errResult
	}
	return order, nil
}

func OrdersProductsFull(ctx context.Context, repo repository.Querier) ([]*repository.OrdersProductsFullRow, error) {
	orders, err := repo.OrdersProductsFull(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func OrderRecount(ctx context.Context, repo repository.Querier, id int) (*repository.ShopOrder, error) {
	_, errID := OrderByID(ctx, repo, id)
	if errID != nil {
		return nil, errID
	}
	errUpdate := repo.OrderUpdateTotal(ctx, int32(id)) //nolint: gosec
	if errUpdate != nil {
		return nil, errUpdate
	}
	order, err := OrderByID(ctx, repo, id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func OrderByID(ctx context.Context, repo repository.Querier, id int) (*repository.ShopOrder, error) {
	order, err := repo.OrderByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("don't found oder by id=%d: %w", id, err)
	}
	return order, nil
}

func OrderDelete(ctx context.Context, repo repository.Querier, idStr string) error {
	id, errConv := strconv.Atoi(idStr)
	if errConv != nil {
		return errConv
	}

	_, errID := OrderByID(ctx, repo, id)
	if errID != nil {
		return errID
	}

	err := repo.OrderDelete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
