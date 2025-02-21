package model

// ProductsResponse represents a response for the GetProducts handler.
type ProductsResponse struct {
	Error    string     `json:"error,omitempty"`
	Products []*Product `json:"products,omitempty"`
}

// PurchaseItem represents an item to be purchased using the PurchaseProducts handler.
type PurchaseItem struct {
	SKU      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

// PurchaseRequest represents a request to purchase a list of type PurchaseItem.
type PurchaseRequest struct {
	Items []PurchaseItem `json:"items"`
}

// PurchaseResponse represents a response for the PurchaseProducts handler.
type PurchaseResponse struct {
	Error           string     `json:"error,omitempty"`
	TotalPrice      float64    `json:"total_price,omitempty"`
	UpdatedProducts []*Product `json:"updated_products,omitempty"`
}
