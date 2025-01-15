package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/m-golang/weather-app/internal/handler"
	"github.com/m-golang/weather-app/internal/helpers"
	"github.com/m-golang/weather-app/internal/middleware"
)

// Router sets up all the routes for the application.
func Router(r *gin.Engine) {
	// Use the middleware to recover from any panic during the request handling.
	r.Use(middleware.RecoverPanic())
	// Use the middleware to add secure headers to all responses.
	r.Use(middleware.SecureHeaders())

	// Set up a fallback route for handling requests that do not match any defined route.
	// If no route is found, it will trigger the `BadRequest` helper and return a 400 status with an error message.
	r.NoRoute(helpers.BadRequest)

	// Define the route for fetching weather data for a specific location.
	// The ":location" is a dynamic URL parameter that will represent the location (e.g., /London).
	// The `WeatherHandler` will be called to process the request and send back the weather data for the location.
	r.GET("/:location", handler.WeatherHandler)
}
