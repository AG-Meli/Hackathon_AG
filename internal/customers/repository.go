package customers

import (
	"database/sql"
	"fmt"
	"github.com/AG-Meli/Hackathon_AG/internal/domain"
	"github.com/AG-Meli/Hackathon_AG/pkg/store"
	"github.com/gin-gonic/gin"
	"strings"
)

var (
	bulkInsertCustomer = "INSERT INTO customers (`id`, `last_name`, `first_name`, `condition`) VALUES %s"
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
	db := r.db
	stmt, err := db.Prepare("INSERT INTO customers (`last_name`, `first_name`, `condition`) VALUES (?, ?, ?)")
	if err != nil {
		return domain.Customer{}, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(lastName, firstName, condition)
	if err != nil {
		return domain.Customer{}, err
	}
	insertedId, err := result.LastInsertId()
	if err != nil {
		return domain.Customer{}, err
	}
	return domain.Customer{ID: int(insertedId), LastName: lastName, FirstName: firstName, Condition: condition}, nil
}

func (r repository) Update(ctx *gin.Context, customerID int, lastName, firstName, condition string) (domain.Customer, error) {
	db := r.db
	stmt, err :=
		db.Prepare("UPDATE customers SET `last_name` = ?, `first_name` = ?, `condition` = ? WHERE `id` = ?;")
	if err != nil {
		return domain.Customer{}, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(lastName, firstName, condition, customerID)
	if err != nil {
		return domain.Customer{}, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return domain.Customer{}, err
	}
	if rowsAffected == 0 {
		return domain.Customer{}, fmt.Errorf("No rows affected")
	}
	return domain.Customer{ID: customerID, LastName: lastName, FirstName: firstName, Condition: condition}, nil
}

func (r repository) Get(ctx *gin.Context, customerID int) (domain.Customer, error) {
	var customer domain.Customer
	db := r.db
	rows, err := db.Query("SELECT `id`, `last_name`, `first_name`, `condition` FROM Customers WHERE `id` = ?;",
		customerID)
	if err != nil {
		return domain.Customer{}, err
	}
	for rows.Next() {
		if err := rows.Scan(&customer.ID, &customer.LastName, &customer.FirstName, &customer.Condition); err != nil {
			return domain.Customer{}, err
		}
	}
	return customer, nil
}

func (r repository) BulkInsert(ctx *gin.Context, customers []domain.Customer) ([]domain.Customer, error) {
	valueStrings := []string{}
	valueArgs := []interface{}{}
	for _, customer := range customers {
		valueStrings = append(valueStrings, "(?, ?, ?, ?)")
		valueArgs = append(valueArgs, customer.ID)
		valueArgs = append(valueArgs, customer.LastName)
		valueArgs = append(valueArgs, customer.FirstName)
		valueArgs = append(valueArgs, customer.Condition)
	}
	bulkInsertCustomer = fmt.Sprintf(bulkInsertCustomer, strings.Join(valueStrings, ","))
	tx, err := r.db.Begin()
	if err != nil {
		return []domain.Customer{}, err
	}
	_, err = tx.Exec(bulkInsertCustomer, valueArgs...)
	if err != nil {
		tx.Rollback()
		return []domain.Customer{}, err
	}
	err = tx.Commit()
	if err != nil {
		return []domain.Customer{}, err
	}
	resetValues()
	return customers, nil
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

func resetValues() {
	bulkInsertCustomer = "INSERT INTO customers (`id`, `last_name`, `first_name`, `condition`) VALUES %s"
}
