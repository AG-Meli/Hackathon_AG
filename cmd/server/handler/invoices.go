package handler

import (
	"github.com/AG-Meli/Hackathon_AG/internal/invoices"
	"github.com/AG-Meli/Hackathon_AG/pkg/web"
	"github.com/gin-gonic/gin"
	"strconv"
)

type HandlerInvoice struct {
	service invoices.Service
}

func NewHandlerInvoice(service invoices.Service) HandlerInvoice {
	return HandlerInvoice{
		service: service,
	}
}

func (i HandlerInvoice) Insert() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestInvoice web.Invoice
		err := ctx.ShouldBind(&requestInvoice)
		if err != nil {
			web.CheckResponse(ctx, web.Response{StatusCode: 400, Content: "Invalid params"})
			return
		}
		response := i.service.Insert(ctx, requestInvoice.DateTime, requestInvoice.CustomerID, requestInvoice.Total)
		web.CheckResponse(ctx, response)
		return
	}
}

func (i HandlerInvoice) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("ID")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			web.CheckResponse(ctx, web.Response{StatusCode: 400, Content: "Invalid ID format"})
			return
		}
		var requestInvoice web.Invoice
		err = ctx.ShouldBind(&requestInvoice)
		if err != nil {
			web.CheckResponse(ctx, web.Response{StatusCode: 400, Content: "Invalid params"})
			return
		}
		response := i.service.Update(ctx, id, requestInvoice.DateTime, requestInvoice.CustomerID, requestInvoice.Total)
		web.CheckResponse(ctx, response)
		return
	}
}

func (i HandlerInvoice) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("ID")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			web.CheckResponse(ctx, web.Response{StatusCode: 400, Content: "Invalid ID format"})
			return
		}
		response := i.service.Get(ctx, id)
		web.CheckResponse(ctx, response)
		return
	}
}

func (i HandlerInvoice) Restore() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fileContent := ctx.Param("fileContent")
		response := i.service.Restore(ctx, fileContent)
		web.CheckResponse(ctx, response)
		return
	}
}
