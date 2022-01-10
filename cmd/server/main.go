package main

import (
	"database/sql"
	"github.com/AG-Meli/Hackathon_AG/cmd/server/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "aghione:aghione90@/hackathon")
	if err != nil {
		panic(err)
	}

	engine := gin.Default()
	router := routes.NewRouter(engine, db)
	router.MapRoutes()
	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}
