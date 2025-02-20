package model

type ProductsResponse struct {
	Error    string     `json:"error,omitempty"`
	Products []*Product `json:"products,omitempty"`
}

type PurchaseItem struct {
	SKU      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

type PurchaseRequest struct {
	Items []PurchaseItem `json:"items"`
}

type PurchaseResponse struct {
	Error           string     `json:"error,omitempty"`
	TotalPrice      float64    `json:"total_price,omitempty"`
	UpdatedProducts []*Product `json:"updated_products,omitempty"`
}
