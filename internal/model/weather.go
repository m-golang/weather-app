package model


// Weather represents the data structure for weather information
type Weather struct {
	Location struct {
		Name           string `json:"name"`
		Country        string `json:"country"`
		LocaltimeEpoch int64  `json:"localtime_epoch"`
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
}