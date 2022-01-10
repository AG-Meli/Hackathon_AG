package sales

import "github.com/AG-Meli/Hackathon_AG/internal/domain"

type Service interface {
	Insert(invoiceID, productID int, quantity float64) (domain.Sale, error)
	Update(saleID, invoiceID, productID int, quantity float64) (domain.Sale, error)
	Get(invoiceID int) (domain.Sale, error)
	Restore(data string) ([]domain.Sale, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s service) Insert(invoiceID, productID int, quantity float64) (domain.Sale, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Update(saleID, invoiceID, productID int, quantity float64) (domain.Sale, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Get(invoiceID int) (domain.Sale, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Restore(data string) ([]domain.Sale, error) {
	//TODO implement me
	panic("implement me")
}
