package model

type Product struct {
	SKU       string  `db:"sku"`
	Name      string  `db:"name"`
	Price     float64 `db:"price"`
	Inventory int     `db:"inventory"`
}
