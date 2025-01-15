package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-golang/weather-app/internal/helpers"
	"github.com/m-golang/weather-app/internal/service"
)

// WeatherHandler handles incoming HTTP requests for weather data based on a location.
// It retrieves weather data for the provided city and returns the data as a JSON response.
func WeatherHandler(c *gin.Context) {
	// Retrieve the "location" parameter from the URL.
	// This will be the city for which weather data is requested.
	city := c.Param("location")

	// If the city parameter is missing or empty, respond with a bad request status.
	if city == "" {
		helpers.BadRequest(c)
		return
	}

	// Call the service to fetch weather data for the city, requesting a 7-day forecast.
	weatherData, err := service.FetchWeatherData(city, 7)
	if err != nil {
		// Handle specific error related to unexpected end of JSON input.
		if errors.Is(err, helpers.ErrUnexpectedEndOfJSONInput) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "No matching location found.",
			})
		} else {
			// For other errors, respond with an internal server error.
			// The error message here can be generic to prevent leaking internal details.
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve weather data. Please try again later.",
			})
		}
		return
	}

	// If no error occurred, send back a JSON response with the weather data.
	c.JSON(http.StatusOK, gin.H{
		"city": city,
		"weatherData":weatherData,
	})

}
