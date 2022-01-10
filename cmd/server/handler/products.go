package handler

import (
	"github.com/AG-Meli/Hackathon_AG/internal/products"
	"github.com/gin-gonic/gin"
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
	return func(c *gin.Context) {

	}
}

func (p HandlerProduct) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (p HandlerProduct) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (p HandlerProduct) Restore() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
