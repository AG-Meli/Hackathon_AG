package handler

import (
	"github.com/AG-Meli/Hackathon_AG/internal/sales"
	"github.com/AG-Meli/Hackathon_AG/pkg/web"
	"github.com/gin-gonic/gin"
	"strconv"
)

type HandlerSale struct {
	service sales.Service
}

func NewHandlerSale(service sales.Service) HandlerSale {
	return HandlerSale{
		service: service,
	}
}

func (s HandlerSale) Insert() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestSale web.Sale
		err := ctx.ShouldBind(&requestSale)
		if err != nil {
			web.CheckResponse(ctx, web.Response{StatusCode: 400, Content: "Invalid params"})
			return
		}
		response := s.service.Insert(ctx, requestSale.InvoiceID, requestSale.ProductID, requestSale.Quantity)
		web.CheckResponse(ctx, response)
		return
	}
}

func (s HandlerSale) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("ID")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			web.CheckResponse(ctx, web.Response{StatusCode: 400, Content: "Invalid ID format"})
			return
		}
		var requestSale web.Sale
		err = ctx.ShouldBind(&requestSale)
		if err != nil {
			web.CheckResponse(ctx, web.Response{StatusCode: 400, Content: "Invalid params"})
			return
		}
		response := s.service.Update(ctx, id, requestSale.InvoiceID, requestSale.ProductID, requestSale.Quantity)
		web.CheckResponse(ctx, response)
		return
	}
}

func (s HandlerSale) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("ID")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			web.CheckResponse(ctx, web.Response{StatusCode: 400, Content: "Invalid ID format"})
			return
		}
		response := s.service.Get(ctx, id)
		web.CheckResponse(ctx, response)
		return
	}
}

func (s HandlerSale) Restore() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fileContent := ctx.Param("fileContent")
		response := s.service.Restore(ctx, fileContent)
		web.CheckResponse(ctx, response)
		return
	}
}
