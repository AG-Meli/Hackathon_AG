package handler

import (
	"github.com/AG-Meli/Hackathon_AG/internal/customers"
	"github.com/gin-gonic/gin"
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
	return func(c *gin.Context) {

	}
}

func (c HandlerCustomer) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (c HandlerCustomer) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (c HandlerCustomer) Restore() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
