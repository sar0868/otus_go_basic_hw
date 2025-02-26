package trx

type Product struct {
	Name     string  `db:"name" json:"name"`
	Quantity float64 `db:"quantity" json:"quantity"`
}

type CreateOrderParams struct {
	User     string `db:"name" json:"user"`
	Products []Product
}
