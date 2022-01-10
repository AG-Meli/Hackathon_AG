package customers

import (
	"github.com/AG-Meli/Hackathon_AG/internal/domain"
	"github.com/AG-Meli/Hackathon_AG/pkg/web"
	"github.com/gin-gonic/gin"
	"strings"
)

type Service interface {
	Insert(ctx *gin.Context, lastName, firstName, condition string) web.Response
	Update(ctx *gin.Context, customerID int, lastName, firstName, condition string) web.Response
	Get(ctx *gin.Context, customerID int) web.Response
	Restore(ctx *gin.Context, location, rowDivider, columnDivider string) web.Response
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s service) Insert(ctx *gin.Context, lastName, firstName, condition string) web.Response {
	if !domain.ValidateCustomerAttributes(lastName, firstName, condition) {
		return web.Response{StatusCode: 400, Content: "Invalid attributes values"}
	}
	customer, err := s.repository.Insert(ctx, lastName, firstName, condition)
	if err != nil {
		return web.Response{StatusCode: 500, Content: err.Error()}
	}
	return web.Response{StatusCode: 201, Content: customer}
}

func (s service) Update(ctx *gin.Context, customerID int, lastName, firstName, condition string) web.Response {
	if !domain.ValidateCustomerAttributes(lastName, firstName, condition) {
		return web.Response{StatusCode: 400, Content: "Invalid attributes values"}
	}
	customer, err := s.repository.Update(ctx, customerID, lastName, firstName, condition)
	if err != nil {
		return web.Response{StatusCode: 500, Content: err.Error()}
	}
	return web.Response{StatusCode: 200, Content: customer}
}

func (s service) Get(ctx *gin.Context, customerID int) web.Response {
	if customerID <= 0 {
		return web.Response{StatusCode: 400, Content: "Invalid ID"}
	}
	customer, err := s.repository.Get(ctx, customerID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return web.Response{StatusCode: 404, Content: "Customer ID not found"}
		}
		return web.Response{StatusCode: 500, Content: err.Error()}
	}
	return web.Response{StatusCode: 200, Content: customer}
}

func (s service) Restore(ctx *gin.Context, location, rowDivider, columnDivider string) web.Response {
	if location == "" || rowDivider == "" || columnDivider == "" {
		return web.Response{400, "Invalid params"}
	}
	customers, err := s.repository.Restore(ctx, location, rowDivider, columnDivider)
	if err != nil {
		return web.Response{StatusCode: 500, Content: err.Error()}
	}
	return web.Response{StatusCode: 201, Content: customers}
}
