package main

import (
	"restapi-golang/config"
	"restapi-golang/handlers"
	"restapi-golang/repositories"
	"restapi-golang/routes"
	"restapi-golang/services"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.ConnectDB()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	mhsRepo := repositories.NewMhsRepository(config.DB)
	mhsService := services.NewMhsService(mhsRepo)
	mhsHandler := handlers.NewMhsHandler(mhsService)
	dashboardMhsRepo := repositories.NewDashboardMhsRepository(config.DB)
	dashboardMhsservice := services.NewDashboardMhsService(dashboardMhsRepo)
	dashboardMhsHandler := handlers.NewDashboardMhsHandler(dashboardMhsservice)

	routes.SetUpRoutes(router, mhsHandler, dashboardMhsHandler)

	router.Run(":8080")
}
