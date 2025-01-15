package helpers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrAPIKeyNotFound is an error that is returned when the weather API key is not set in the environment variables.
var ErrAPIKeyNotFound = errors.New("weatherapi key not set in environment variables")

// ErrUnexpectedEndOfJSONInput is an error that is returned when the weather API response contains incomplete JSON data.
var ErrUnexpectedEndOfJSONInput = errors.New("unexpected end of JSON input")

// BadRequest sends a 400 Bad Request response when the 'location' parameter is missing or invalid.
func BadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Missing or invalid 'location' parameter",
	})
}
