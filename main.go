package main

import (
	"log"
	"strings"
	"time"

	"github.com/ShiranaiZo/experiment-golang/app/controllers"
	"github.com/ShiranaiZo/experiment-golang/app/repositories"
	"github.com/ShiranaiZo/experiment-golang/app/services"
	"github.com/ShiranaiZo/experiment-golang/config"
	"github.com/ShiranaiZo/experiment-golang/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from the .env file.
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Initialize the config env and database connection.
	config.InitMainConfig()
	db, err := config.InitDatabase()

	if err != nil {
		panic("Failed to initialize database: " + err.Error())
	}

	// Create a new Gin router with default middleware (logger and recovery).
	router := gin.Default()

	// Configure CORS (Cross-Origin Resource Sharing) settings.
	corsConfig := cors.Config{
		AllowOrigins:     strings.Split(config.MainConfig.CORSAllowedOrigins, ","),
		AllowMethods:     strings.Split(config.MainConfig.CORSAllowedMethods, ","),
		AllowHeaders:     strings.Split(config.MainConfig.CORSAllowedHeaders, ","),
		AllowCredentials: config.MainConfig.CORSAllowCredentials,
		ExposeHeaders:    strings.Split(config.MainConfig.CORSExposeHeaders, ","),
		MaxAge:           time.Duration(config.MainConfig.CORSMaxAge) * time.Hour,
	}

	// Apply the CORS middleware to the router.
	router.Use(cors.New(corsConfig))

	// Initialize seeders, repositories, services, and controllers.
	// seeders.NewSeederRegistry(config.DB).Run()
	repository := repositories.NewRepositoryRegistry(db)
	service := services.NewServiceRegistry(repository)
	controller := controllers.NewControllerRegistry(service)

	// Define application routes and associate them with their respective handlers.
	routes.InitRoutes(router, controller)

	// Start the HTTP server on the specified port.
	err = router.Run(":" + config.MainConfig.AppPORT)

	if err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
