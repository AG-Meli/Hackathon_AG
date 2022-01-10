package routes

import (
	"database/sql"
	"github.com/AG-Meli/Hackathon_AG/cmd/server/handler"
	"github.com/AG-Meli/Hackathon_AG/internal/customers"
	"github.com/AG-Meli/Hackathon_AG/internal/invoices"
	"github.com/AG-Meli/Hackathon_AG/internal/products"
	"github.com/AG-Meli/Hackathon_AG/internal/sales"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	engine *gin.Engine
	rg     *gin.RouterGroup
	db     *sql.DB
}

func NewRouter(engine *gin.Engine, db *sql.DB) Router {
	return &router{
		engine: engine,
		db:     db,
	}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildProductRoutes()
	r.buildSalesRoutes()
	r.buildCustomersRoutes()
	r.buildInvoicesRoutes()
}

func (r *router) setGroup() {
	r.rg = r.engine.Group("/api/v1")
}

func (r *router) buildProductRoutes() {
	repo := products.NewRepository(r.db)
	service := products.NewService(repo)
	handler := handler.NewHandlerProduct(service)
	productsGroup := r.rg.Group("/products")
	{
		productsGroup.POST("/", handler.Insert())
		productsGroup.PUT("/:ID", handler.Update())
		productsGroup.GET("/:ID", handler.Get())
		productsGroup.POST("/restore/", handler.Restore())
	}
}

func (r *router) buildSalesRoutes() {
	repo := sales.NewRepository(r.db)
	service := sales.NewService(repo)
	handler := handler.NewHandlerSale(service)
	saleGroup := r.rg.Group("/sales")
	{
		saleGroup.POST("/", handler.Insert())
		saleGroup.PUT("/", handler.Update())
		saleGroup.GET("/", handler.Get())
		saleGroup.POST("/restore/", handler.Restore())
	}
}

func (r *router) buildCustomersRoutes() {
	repo := customers.NewRepository(r.db)
	service := customers.NewService(repo)
	handler := handler.NewHandlerCustomer(service)
	customersGroup := r.rg.Group("/customers")
	{
		customersGroup.POST("/", handler.Insert())
		customersGroup.PUT("/:ID", handler.Update())
		customersGroup.GET("/:ID", handler.Get())
		customersGroup.POST("/restore/", handler.Restore())
	}
}

func (r *router) buildInvoicesRoutes() {
	repo := invoices.NewRepository(r.db)
	service := invoices.NewService(repo)
	handler := handler.NewHandlerInvoice(service)
	invoicesGroup := r.rg.Group("/invoices")
	{
		invoicesGroup.POST("/", handler.Insert())
		invoicesGroup.PUT("/", handler.Update())
		invoicesGroup.GET("/", handler.Get())
		invoicesGroup.POST("/restore/", handler.Restore())
	}
}
