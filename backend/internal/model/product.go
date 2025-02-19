package model

type Product struct {
	SKU       string  `db:"sku" json:"sku"`
	Name      string  `db:"name" json:"name"`
	Price     float64 `db:"price" json:"price"`
	Inventory int     `db:"inventory" json:"inventory"`
}
