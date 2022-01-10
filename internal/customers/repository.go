package customers

import (
	"database/sql"
	"github.com/AG-Meli/Hackathon_AG/internal/domain"
	"github.com/AG-Meli/Hackathon_AG/pkg/store"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	Insert(ctx *gin.Context, lastName, firstName, condition string) (domain.Customer, error)
	Update(ctx *gin.Context, customerID int, lastName, firstName, condition string) (domain.Customer, error)
	Get(ctx *gin.Context, customerID int) (domain.Customer, error)
	Restore(ctx *gin.Context, fileLocation, rowDivider, columnDivider string) ([]domain.Customer, error)
	BulkInsert(ctx *gin.Context, customers []domain.Customer) ([]domain.Customer, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) Insert(ctx *gin.Context, lastName, firstName, condition string) (domain.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Update(ctx *gin.Context, customerID int, lastName, firstName, condition string) (domain.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Get(ctx *gin.Context, customerID int) (domain.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) BulkInsert(ctx *gin.Context, customers []domain.Customer) ([]domain.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Restore(ctx *gin.Context, fileLocation, rowDivider, columnDivider string) ([]domain.Customer, error) {
	fileManager := store.NewFileManager(fileLocation, rowDivider, columnDivider)
	stringContent, err := fileManager.ReadFile()
	if err != nil {
		return []domain.Customer{}, nil
	}
	customers, err := fileManager.CovertToCustomers(stringContent)
	if err != nil {
		return []domain.Customer{}, err
	}
	customers, err = r.BulkInsert(ctx, customers)
	if err != nil {
		return []domain.Customer{}, err
	}
	return customers, nil
}
