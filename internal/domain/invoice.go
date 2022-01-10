package domain

type Invoice struct {
	ID         int     `json:"id"`
	DateTime   string  `json:"date_time"`
	CustomerID int     `json:"customer_id"`
	Total      float64 `json:"total"`
}
