package main

import (
	"assignment_2/config"
	"assignment_2/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.POST("/createOrders", inDB.CreateOrders)
	router.GET("/orders", inDB.GetOrders)
	router.PUT("/orders/:id", inDB.UpdateOrders)
	router.DELETE("/orders/:id", inDB.DeleteOrders)

	router.Run(":3300")

}
