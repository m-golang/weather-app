package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/m-golang/weather-app/internal/config"
	"github.com/m-golang/weather-app/internal/helpers"
	"github.com/m-golang/weather-app/internal/model"
)

// Constant for the environment variable key name that holds the weather API key
const weatherApiKeyName = "WEATHER_API_KEY"

// FetchWeatherData retrieves weather data for a given location and number of forecast days
func FetchWeatherData(location string, days int) (*model.Weather, error) {
	// Load the weather API key from the environment variable
	apiKey, err := config.LoadAPIKey(weatherApiKeyName)
	if err != nil {
		return nil, err
	}

	// Construct the API request URL with the location and the number of days for the forecast
	url := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=%d&aqi=no&alerts=no", apiKey, location, days)

	// Call the helper function to make the HTTP request to the weather API
	responseBody, err := RequestToWeatherApi(url)
	if err != nil {
		return nil, err
	}

	// Declare a variable to hold the parsed weather data
	var weatherData model.Weather

	// Parse the response body (JSON) into the weatherData variable
	err = json.Unmarshal(responseBody, &weatherData)
	if err != nil {
		// Check if the error is related to invalid or incomplete JSON syntax
		if _, ok := err.(*json.SyntaxError); ok {
			// Return a specific error if the JSON parsing fails due to unexpected end of JSON input
			return nil, helpers.ErrUnexpectedEndOfJSONInput
		}
		// Return a general error if JSON unmarshaling fails for any other reason
		return nil, err
	}

	// Return the parsed weather data and nil for the error if everything succeeds
	return &weatherData, nil
}

// RequestToWeatherApi makes an HTTP GET request to the Weather API and returns the response body
func RequestToWeatherApi(url string) ([]byte, error) {
	// Make a GET request to the provided URL
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// Ensure that the response body is closed once the function finishes executing
	defer response.Body.Close()

	// Check if the HTTP response status code is 200 OK, if not, return an error
	if response.StatusCode != http.StatusOK {
		return nil, err
	}

	// Read the entire response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Return the response body as a byte slice
	return body, nil
}
