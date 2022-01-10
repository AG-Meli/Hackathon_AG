package products

import "github.com/AG-Meli/Hackathon_AG/internal/domain"

type Service interface {
	Insert(dateTime string, customerID int, total float64) (domain.Product, error)
	Update(id int, dateTime string, customerID int, total float64) (domain.Product, error)
	Get(id int) (domain.Product, error)
	RestoreData(data string) ([]domain.Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s service) Insert(dateTime string, customerID int, total float64) (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Update(id int, dateTime string, customerID int, total float64) (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Get(id int) (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) RestoreData(data string) ([]domain.Product, error) {
	//TODO implement me
	panic("implement me")
}
