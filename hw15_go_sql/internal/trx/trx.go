package trx

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sar0868/otus_go_basic_hw/hw15_go_sql/internal/repository"
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
		var quant pgtype.Numeric
		quant.Scan(params.Products[product].Quantity)
		paramProduct := repository.ProductAddOrderParams{
			OrderID:  int32(orderID), //nolint: gosec
			Name:     params.Products[product].Name,
			Quantity: quant,
		}
		errAddProduct := repo.ProductAddOrder(ctx, paramProduct)
		if errAddProduct != nil {
			return nil, errAddProduct
		}
		errOrdUp := repo.OrderUpdateTotal(ctx, int32(orderID)) //nolint: gosec
		if errOrdUp != nil {
			return nil, errOrdUp
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
