package customers

import (
	"database/sql"
	"github.com/AG-Meli/Hackathon_AG/internal/domain"
)

type Repository interface {
	Insert(lastName, firstName, condition string) (domain.Customer, error)
	Update(customerID int, firstName, condition string) (domain.Customer, error)
	Get(customerID int) (domain.Customer, error)
	Restore(data string) ([]domain.Customer, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) Insert(lastName, firstName, condition string) (domain.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Update(customerID int, firstName, condition string) (domain.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Get(customerID int) (domain.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Restore(data string) ([]domain.Customer, error) {
	//TODO implement me
	panic("implement me")
}
