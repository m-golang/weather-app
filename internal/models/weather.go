package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Weather represents the data structure for weather information
type Weather struct {
	Location struct {
		Name           string `json:"name"`
		Country        string `json:"country"`
		LocaltimeEpoch int64  `json:"localtime_epoch"`
		FormattedTime  string
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
		} `json:"condition"`
		Humidity   int     `json:"humidity"`
		WindKph    float64 `json:"wind_kph"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		Precip_mm  float64 `json:"precip_mm"`
		VisKm      float64 `json:"vis_km"`
		Cloud      int     `json:"cloud"`
	} `json:"current"`
	Forecast struct {
		ForecastDays []ForecastDay `json:"forecastday"`
	} `json:"forecast"`
}

// ForecastDay holds information for a single day's forecast
type ForecastDay struct {
	DateEpoch int64 `json:"date_epoch"`
	Day       struct {
		MaxTempC          float64 `json:"maxtemp_c"`
		MinTempC          float64 `json:"mintemp_c"`
		DailyWillItRain   int     `json:"daily_will_it_rain"`
		DailyChanceOfRain int     `json:"daily_chance_of_rain"`
		DailyWillItSnow   int     `json:"daily_will_it_snow"`
		DailyChanceOfSnow int     `json:"daily_chance_of_snow"`
		Condition         struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
		} `json:"condition"`
	} `json:"day"`
	FormattedTime string
	DayName       string
}

// GetWeatherData fetches weather data from the API
func GetWeatherData(location string, days int) (*Weather, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("Weather API key not set in environment variables")
	}

	url := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=%d&aqi=no&alerts=no", apiKey, location, days)

	// Send a GET request to the weather API
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// If the response is not OK, return an error
	if resp.StatusCode != 200 {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherData Weather

	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return nil, err
	}

	return &weatherData, nil
}

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
