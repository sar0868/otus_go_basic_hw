package service

import (
	"context"
	"fmt"

	"github.com/sar0868/otus_go_basic_hw/hw16_docker/internal/repository"
)

func Products(ctx context.Context, repo repository.Querier) ([]*repository.ShopProduct, error) {
	products, err := repo.Products(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func ProductByName(ctx context.Context, repo repository.Querier, name string) (*repository.ShopProduct, error) {
	product, err := repo.ProductGetByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("don't found product for name=%s", name)
	}
	return product, nil
}

func ProductCreate(ctx context.Context, repo repository.Querier, newProduct repository.ProductCreateParams) (*repository.ShopProduct, error) { //nolint: lll
	id, err := repo.ProductCreate(ctx, newProduct)
	if err != nil {
		return nil, fmt.Errorf("don't create product: %w", err)
	}
	product, _ := repo.ProductGetById(ctx, id)
	return product, nil
}

func ProductUpdate(ctx context.Context, repo repository.Querier, newProduct repository.ProductUpdateParams) (*repository.ShopProduct, error) { //nolint: lll
	if err := repo.ProductUpdate(ctx, newProduct); err != nil {
		return nil, err
	}
	product, err := ProductByName(ctx, repo, newProduct.Name)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func DeleteProduct(ctx context.Context, repo repository.Querier, name string) error {
	_, errID := repo.ProductID(ctx, name)
	if errID != nil {
		return fmt.Errorf("error don't found product: %w", errID)
	}
	err := repo.ProductDelete(ctx, name)
	if err != nil {
		return fmt.Errorf("error delete product: %w", err)
	}
	return nil
}

func ProductByID(ctx context.Context, repo repository.Querier, id int) (*repository.ShopProduct, error) {
	product, err := repo.ProductGetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("don't found product by id=%d", id)
	}
	return product, nil
}

func ProductsGetRangePrice(ctx context.Context, repo repository.Querier, paramsPrice repository.ProductGetRangePriceParams) ([]*repository.ShopProduct, error) { //nolint: lll
	products, err := repo.ProductGetRangePrice(ctx, paramsPrice)
	if err != nil {
		return nil, err
	}
	return products, nil
}
