package handler

import (
	"github.com/AG-Meli/Hackathon_AG/internal/products"
	"github.com/AG-Meli/Hackathon_AG/pkg/web"
	"github.com/gin-gonic/gin"
	"strconv"
)

type HandlerProduct struct {
	service products.Service
}

func NewHandlerProduct(service products.Service) HandlerProduct {
	return HandlerProduct{
		service: service,
	}
}

func (p HandlerProduct) Insert() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestProduct web.Product
		err := ctx.ShouldBind(&requestProduct)
		if err != nil {
			web.CheckResponse(ctx, web.Response{StatusCode: 400, Content: "Invalid params"})
			return
		}
		response := p.service.Insert(ctx, requestProduct.Description, requestProduct.Price)
		web.CheckResponse(ctx, response)
		return
	}
}

func (p HandlerProduct) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("ID")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			web.CheckResponse(ctx, web.Response{StatusCode: 400, Content: "Invalid ID format"})
			return
		}
		var requestProduct web.Product
		err = ctx.ShouldBind(&requestProduct)
		if err != nil {
			web.CheckResponse(ctx, web.Response{StatusCode: 400, Content: "Invalid params"})
			return
		}
		response := p.service.Update(ctx, id, requestProduct.Description, requestProduct.Price)
		web.CheckResponse(ctx, response)
		return
	}
}

func (p HandlerProduct) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("ID")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			web.CheckResponse(ctx, web.Response{StatusCode: 400, Content: "Invalid ID format"})
			return
		}
		response := p.service.Get(ctx, id)
		web.CheckResponse(ctx, response)
		return
	}
}

func (p HandlerProduct) Restore() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fileContent := ctx.Param("fileContent")
		response := p.service.Restore(ctx, fileContent)
		web.CheckResponse(ctx, response)
		return
	}
}
