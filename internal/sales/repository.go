package sales

import (
	"database/sql"
	"github.com/AG-Meli/Hackathon_AG/internal/domain"
)

type Repository interface {
	Insert(invoiceID, productID int, quantity float64) (domain.Sale, error)
	Update(saleID, invoiceID, productID int, quantity float64) (domain.Sale, error)
	Get(invoiceID int) (domain.Sale, error)
	Restore(data string) ([]domain.Sale, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) Insert(invoiceID, productID int, quantity float64) (domain.Sale, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Update(saleID, invoiceID, productID int, quantity float64) (domain.Sale, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Get(invoiceID int) (domain.Sale, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Restore(data string) ([]domain.Sale, error) {
	//TODO implement me
	panic("implement me")
}
