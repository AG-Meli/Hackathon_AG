package store

import (
	"github.com/AG-Meli/Hackathon_AG/internal/domain"
	"io/ioutil"
	"strconv"
	"strings"
)

type FileManager interface {
	ReadFile() (string, error)
	CovertToCustomers(customersStr string) ([]domain.Customer, error)
	ConvertToSales(salesStr string) ([]domain.Sale, error)
	ConvertToProducts(productsStr string) ([]domain.Product, error)
	ConvertToInvoices(invoicesStr string) ([]domain.Invoice, error)
}

type fileManager struct {
	FileLocation  string
	RowDivider    string
	ColumnDivider string
}

func NewFileManager(fileLocation, rowDivider, columnDivider string) FileManager {
	return &fileManager{
		FileLocation:  fileLocation,
		RowDivider:    rowDivider,
		ColumnDivider: columnDivider,
	}
}

func (f fileManager) ReadFile() (string, error) {
	file, err := ioutil.ReadFile(f.FileLocation)
	if err != nil {
		return "", err
	}
	data := string(file)
	return data, err
}

func (f fileManager) CovertToCustomers(customersStr string) ([]domain.Customer, error) {
	var customers []domain.Customer
	arrayDataStr := strings.Split(customersStr, f.RowDivider)
	for _, value := range arrayDataStr {
		customerParamsStr := strings.Split(value, f.ColumnDivider)
		id, err := strconv.Atoi(customerParamsStr[0])
		if err != nil {
			return []domain.Customer{}, err
		}
		var customer = domain.Customer{
			ID:        id,
			LastName:  customerParamsStr[1],
			FirstName: customerParamsStr[2],
			Condition: customerParamsStr[3],
		}
		if !domain.ValidateCustomerAttributes(customer.LastName, customer.FirstName, customer.Condition) {
			return []domain.Customer{}, nil
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (f fileManager) ConvertToSales(salesStr string) ([]domain.Sale, error) {
	//TODO implement me
	panic("implement me")
}

func (f fileManager) ConvertToProducts(productsStr string) ([]domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (f fileManager) ConvertToInvoices(invoicesStr string) ([]domain.Invoice, error) {
	//TODO implement me
	panic("implement me")
}
