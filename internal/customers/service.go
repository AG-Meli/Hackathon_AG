package customers

import "github.com/AG-Meli/Hackathon_AG/internal/domain"

type Service interface {
	Insert(lastName, firstName, condition string) (domain.Customer, error)
	Update(customerID int, firstName, condition string) (domain.Customer, error)
	Get(customerID int) (domain.Customer, error)
	Restore(data string) ([]domain.Customer, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s service) Insert(lastName, firstName, condition string) (domain.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Update(customerID int, firstName, condition string) (domain.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Get(customerID int) (domain.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Restore(data string) ([]domain.Customer, error) {
	//TODO implement me
	panic("implement me")
}
