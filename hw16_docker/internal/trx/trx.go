package trx

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sar0868/otus_go_basic_hw/hw16_docker/internal/repository"
)

func CreateOrderWithProducts(ctx context.Context, params CreateOrderParams, db *pgxpool.Pool) (*repository.OrderUserRow, error) { //nolint: lll
	var orderID int

	tx, err := db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	dbTx := repository.Queries{}
	repo := dbTx.WithTx(tx)

	orderID, err = repo.OrderCreateByUser(ctx, params.User)
	if err != nil {
		return nil, err
	}
	for product := range params.Products {
		idPr, errProduct := repo.ProductID(ctx, params.Products[product].Name)
		if errProduct != nil {
			return nil, fmt.Errorf("get id product: %w", errProduct)
		}
		paramProduct := repository.OrderProductsCreateParams{
			OrderID:   int32(orderID), //nolint: gosec
			ProductID: int32(idPr),    //nolint: gosec
			Quantity:  params.Products[product].Quantity,
		}

		errAddProduct := repo.OrderProductsCreate(ctx, paramProduct)
		if errAddProduct != nil {
			return nil, fmt.Errorf("product add order :%w", errAddProduct)
		}
		errOrdUp := repo.OrderUpdateTotal(ctx, int32(orderID)) //nolint: gosec
		if errOrdUp != nil {
			return nil, fmt.Errorf("error update total: %w", errOrdUp)
		}
	}
	order, errOrder := repo.OrderUser(ctx, orderID)
	if errOrder != nil {
		return nil, errOrder
	}
	if errCommit := tx.Commit(ctx); errCommit != nil {
		return nil, errCommit
	}

	return order, nil
}

func OrderAddProduct(ctx context.Context, params repository.ProductAddOrderParams, db *pgxpool.Pool) (*repository.OrderUserRow, error) { //nolint: lll
	tx, err := db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	dbTx := repository.Queries{}
	repo := dbTx.WithTx(tx)

	idPr, errProduct := repo.ProductID(ctx, params.Name)
	if errProduct != nil {
		return nil, fmt.Errorf("get id product: %w", errProduct)
	}
	product := repository.OrderProductsCreateParams{
		OrderID:   params.OrderID,
		ProductID: int32(idPr), //nolint: gosec
		Quantity:  params.Quantity,
	}
	if errAddProduct := repo.OrderProductsCreate(ctx, product); errAddProduct != nil {
		return nil, errAddProduct
	}

	errOrdUp := repo.OrderUpdateTotal(ctx, params.OrderID)
	if errOrdUp != nil {
		return nil, fmt.Errorf("error update total: %w", errOrdUp)
	}

	id := int(params.OrderID)
	order, err := repo.OrderUser(ctx, id)
	if err != nil {
		return nil, err
	}

	if errCommit := tx.Commit(ctx); errCommit != nil {
		return nil, errCommit
	}

	return order, nil
}
