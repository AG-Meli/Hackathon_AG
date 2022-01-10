package invoices

import (
	"database/sql"
	"github.com/AG-Meli/Hackathon_AG/internal/domain"
)

type Repository interface {
	Insert(dateTime string, customerID int, total float64) (domain.Invoice, error)
	Update(invoiceID int, dateTime string, customerID int, total float64) (domain.Invoice, error)
	Get(invoiceID int) (domain.Invoice, error)
	Restore(data string) ([]domain.Invoice, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) Insert(dateTime string, customerID int, total float64) (domain.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Update(invoiceID int, dateTime string, customerID int, total float64) (domain.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Get(invoiceID int) (domain.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Restore(data string) ([]domain.Invoice, error) {
	//TODO implement me
	panic("implement me")
}
