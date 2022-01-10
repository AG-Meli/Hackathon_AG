package products

import (
	"github.com/AG-Meli/Hackathon_AG/pkg/web"
	"github.com/gin-gonic/gin"
)

type Service interface {
	Insert(ctx *gin.Context, description string, price float64) web.Response
	Update(ctx *gin.Context, id int, description string, price float64) web.Response
	Get(ctx *gin.Context, id int) web.Response
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

func (s service) Insert(ctx *gin.Context, description string, price float64) web.Response {
	//TODO implement me
	panic("implement me")
}

func (s service) Update(ctx *gin.Context, id int, description string, price float64) web.Response {
	//TODO implement me
	panic("implement me")
}

func (s service) Get(ctx *gin.Context, id int) web.Response {
	//TODO implement me
	panic("implement me")
}

func (s service) Restore(ctx *gin.Context, data string) web.Response {
	//TODO implement me
	panic("implement me")
}
