package handler

import (
	"github.com/AG-Meli/Hackathon_AG/internal/customers"
	"github.com/AG-Meli/Hackathon_AG/pkg/web"
	"github.com/gin-gonic/gin"
	"strconv"
)

type HandlerCustomer struct {
	service customers.Service
}

func NewHandlerCustomer(service customers.Service) HandlerCustomer {
	return HandlerCustomer{
		service: service,
	}
}

func (c HandlerCustomer) Insert() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestCustomer web.Customer
		err := ctx.ShouldBind(&requestCustomer)
		if err != nil {
			web.CheckResponse(ctx, web.Response{StatusCode: 400, Content: "Invalid params"})
			return
		}
		response := c.service.Insert(ctx, requestCustomer.LastName, requestCustomer.FirstName, requestCustomer.Condition)
		web.CheckResponse(ctx, response)
		return
	}
}

func (c HandlerCustomer) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("ID")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			web.CheckResponse(ctx, web.Response{StatusCode: 400, Content: "Invalid ID format"})
			return
		}
		var requestCustomer web.Customer
		err = ctx.ShouldBind(&requestCustomer)
		if err != nil {
			web.CheckResponse(ctx, web.Response{StatusCode: 400, Content: "Invalid params"})
			return
		}
		response := c.service.Update(ctx, id, requestCustomer.LastName, requestCustomer.FirstName, requestCustomer.Condition)
		web.CheckResponse(ctx, response)
		return
	}
}

func (c HandlerCustomer) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("ID")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			web.CheckResponse(ctx, web.Response{StatusCode: 400, Content: "Invalid ID format"})
			return
		}
		response := c.service.Get(ctx, id)
		web.CheckResponse(ctx, response)
		return
	}
}

func (c HandlerCustomer) Restore() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fileContent := ctx.Param("fileContent")
		response := c.service.Restore(ctx, fileContent)
		web.CheckResponse(ctx, response)
		return
	}
}
