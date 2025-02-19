package model

type ProductsResponse struct {
	Products []*Product `json:"products"`
}

type PurchaseItem struct {
	SKU      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

type PurchaseRequest struct {
	Items []PurchaseItem `json:"items"`
}

type PurchaseResponse struct {
	TotalPrice      float64    `json:"total_price"`
	UpdatedProducts []*Product `json:"updated_products"`
}
