package invoices

import "github.com/AG-Meli/Hackathon_AG/internal/domain"

type Service interface {
	Insert(dateTime string, customerID int, total float64) (domain.Invoice, error)
	Update(invoiceID int, dateTime string, customerID int, total float64) (domain.Invoice, error)
	Get(invoiceID int) (domain.Invoice, error)
	Restore(data string) ([]domain.Invoice, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s service) Insert(dateTime string, customerID int, total float64) (domain.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Update(invoiceID int, dateTime string, customerID int, total float64) (domain.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Get(invoiceID int) (domain.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Restore(data string) ([]domain.Invoice, error) {
	//TODO implement me
	panic("implement me")
}
