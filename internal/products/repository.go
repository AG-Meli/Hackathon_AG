package products

import (
	"database/sql"
	"github.com/AG-Meli/Hackathon_AG/internal/domain"
)

type Repository interface {
	Insert(dateTime string, customerID int, total float64) (domain.Product, error)
	Update(id int, dateTime string, customerID int, total float64) (domain.Product, error)
	Get(id int) (domain.Product, error)
	RestoreData(data string) ([]domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) Insert(dateTime string, customerID int, total float64) (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Update(id int, dateTime string, customerID int, total float64) (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Get(id int) (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) RestoreData(data string) ([]domain.Product, error) {
	//TODO implement me
	panic("implement me")
}
