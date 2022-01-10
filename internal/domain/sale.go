package domain

type Sale struct {
	ID        int     `json:"id"`
	InvoiceID int     `json:"invoice_id"`
	ProductID int     `json:"product_id"`
	Quantity  float64 `json:"quantity"`
}
