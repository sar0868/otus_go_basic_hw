package trx

import "github.com/jackc/pgx/v5/pgtype"

// type Product struct {
// 	Name     string         `db:"name" json:"name"`
// 	Quantity pgtype.Numeric `db:"quantity" json:"quantity"`
// }

// type CreateOrderParams struct {
// 	User     string    `db:"name" json:"user"`
// 	Products []Product `json:"products"`
// }

type CreateOrderParams struct {
	User     string         `db:"name" json:"user"`
	Name     string         `db:"name" json:"product"`
	Quantity pgtype.Numeric `db:"quantity" json:"quantity"`
}
