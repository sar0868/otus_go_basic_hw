package trx

import (
	"context"
	"fmt"

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
		return nil, fmt.Errorf("error create order: %w", err)
	}
	// for product := range params.Products {

	idPr, errProduct := repo.ProductID(ctx, params.Name)
	if errProduct != nil {
		return nil, errProduct
	}

	var quant pgtype.Numeric
	// quant.Scan(params.Products[product].Quantity)
	quant.Scan(params.Quantity)
	paramProduct := repository.OrderProductsCreateParams{
		OrderID: int32(orderID), //nolint: gosec
		// Name:     params.Products[product].Name,
		ProductID: int32(idPr), //nolint: gosec
		Quantity:  quant,
	}

	errAddProduct := repo.OrderProductsCreate(ctx, paramProduct)
	if errAddProduct != nil {
		return nil, fmt.Errorf("product add order :%w", errAddProduct)
	}
	errOrdUp := repo.OrderUpdateTotal(ctx, int32(orderID)) //nolint: gosec
	if errOrdUp != nil {
		return nil, errOrdUp
	}
	// }
	order, errOrder := repo.OrderUser(ctx, orderID)
	if errOrder != nil {
		return nil, errOrder
	}

	if errCommit := tx.Commit(ctx); errCommit != nil {
		return nil, errCommit
	}
	return order, nil
}
