package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/m-golang/weather-app/internal/models"
)

// home handles the root route and redirects to a default city's weather view.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Redirect to the weather page of Tashkent
	city := "tashkent"
	http.Redirect(w, r, fmt.Sprintf("/weather/%s", city), http.StatusSeeOther)
}

// weatherView displays the weather information for a given city.
func (app *application) weatherView(w http.ResponseWriter, r *http.Request) {
	// Extract city name from URL parameters
	params := httprouter.ParamsFromContext(r.Context())
	city := params.ByName("city")

	// If no city is provided, return a 404 error
	if city == "" {
		app.notFound(w)
		return
	}

	// Get weather forecast data for the city (7 day forecast by default)
	Weather, err := models.GetWeatherData(city, 7)
	if err != nil {
		app.serverError(w, fmt.Errorf("could not get weather data for %s: %v", city, err))
		return
	}

	// If weather data is nil, return a 404 error
	if Weather == nil {
		app.notFound(w)
		return
	}

	// If weather data is incomplete, return a 404 error
	if Weather.Location.Name == "" || Weather.Current.Condition.Text == "" {
		app.clientError(w, http.StatusNotFound)
		return
	}

	// Prepare data for rendering the template
	data := app.newTempalteData(r)
	data.Weather = Weather

	// Convert the current local time to a formatted string
	localTime := time.Unix(data.Weather.Location.LocaltimeEpoch, 0)
	formattedTime := localTime.Format("15:04")
	data.Weather.Location.FormattedTime = formattedTime

	// Convert forecast data times to formatted strings
	for i := range data.Weather.Forecast.ForecastDays {
		foreCastTime := time.Unix(data.Weather.Forecast.ForecastDays[i].DateEpoch, 0)
		formattedTimeForecast := foreCastTime.Format("02 Jan")
		dayName := foreCastTime.Format("Mon")

		data.Weather.Forecast.ForecastDays[i].DayName = dayName
		data.Weather.Forecast.ForecastDays[i].FormattedTime = formattedTimeForecast
	}

	// Render the weather view template with the data
	app.render(w, http.StatusOK, "view.html", data)
}
