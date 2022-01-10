package customers

import (
	"github.com/AG-Meli/Hackathon_AG/pkg/web"
	"github.com/gin-gonic/gin"
)

type Service interface {
	Insert(ctx *gin.Context, lastName, firstName, condition string) web.Response
	Update(ctx *gin.Context, customerID int, lastName, firstName, condition string) web.Response
	Get(ctx *gin.Context, customerID int) web.Response
	Restore(ctx *gin.Context, data string) web.Response
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
	//TODO implement me
	panic("implement me")
}

func (s service) Update(ctx *gin.Context, customerID int, lastName, firstName, condition string) web.Response {
	//TODO implement me
	panic("implement me")
}

func (s service) Get(ctx *gin.Context, customerID int) web.Response {
	//TODO implement me
	panic("implement me")
}

func (s service) Restore(ctx *gin.Context, data string) web.Response {
	//TODO implement me
	panic("implement me")
}
