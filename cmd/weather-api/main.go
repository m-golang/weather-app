package main

import (
	"github.com/gin-gonic/gin"
	"github.com/m-golang/weather-app/internal/config"
	"github.com/m-golang/weather-app/internal/routes"
)

// main function initializes the application and starts the server.
func main() {
	// Load environment variables from a .env file
	config.LoadEnvFile()

	// Create a new Gin router instance
	r := gin.Default()

	// Set up all routes for the application
	routes.Router(r)

	// Start the web server on the default port (8080)
	r.Run()
}
