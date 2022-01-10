package web

type Customer struct {
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Condition string `json:"condition"`
}

type Invoice struct {
	DateTime   string  `json:"date_time"`
	CustomerID int     `json:"customer_id"`
	Total      float64 `json:"total"`
}

type Product struct {
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type Sale struct {
	InvoiceID int     `json:"invoice_id"`
	ProductID int     `json:"product_id"`
	Quantity  float64 `json:"quantity"`
}

type FileData struct {
	FileLocation  string `json:"file_location"`
	RowDivider    string `json:"row_divider"`
	ColumnDivider string `json:"column_divider"`
}
