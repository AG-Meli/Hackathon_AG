package handler

import (
	"github.com/AG-Meli/Hackathon_AG/internal/sales"
	"github.com/gin-gonic/gin"
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
	return func(c *gin.Context) {

	}
}

func (s HandlerSale) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (s HandlerSale) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (s HandlerSale) Restore() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
