package handler

import (
	"github.com/AG-Meli/Hackathon_AG/internal/invoices"
	"github.com/gin-gonic/gin"
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
	return func(c *gin.Context) {

	}
}

func (i HandlerInvoice) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (i HandlerInvoice) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (i HandlerInvoice) Restore() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
