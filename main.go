package main

import (
	"elewa/pkg/config"
	"elewa/pkg/middleware"
	"elewa/pkg/routes"
	"os"

	"github.com/gin-gonic/gin"
)

var mongoURI string

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)
	router.Use(middleware.Authentication)

	routes.BookRoutes(router)
	routes.CategoryRoutes(router)
	routes.OrderRoutes(router)
	routes.InvoiceRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + port)

	mongoURI = os.Getenv("MONGODB_URI")
	client := config.DbInstance()
	config.Client = client

}
