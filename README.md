# Weather API (First Project)

A simple weather API built with Go (Gin framework) that fetches weather data based on location and provides a 7-day forecast.


**This is my first Go project, and I am learning how to build web applications.**

## Features

- Fetch current weather and 7-day forecast.
- Handles location input via URL parameter.
- Error handling for invalid location or API issues.

## Getting Started

### Prerequisites
- Go 1.16+
- Weather API Key from [weatherapi.com](https://weatherapi.com)

### Installation

1. Clone the repo:
   ```bash
   git clone https://github.com/yourusername/weather-api.git
   cd weather-api
2. Install dependencies:
   ```bash
   go mod tidy
3. Set up your .env file with your Weather API key:
   ```bash
   WEATHER_API_KEY=your-weather-api-key
4. Run the server:
   ```bash
   go run main.go
5. Access the weather data by navigating to:
   ```bash
   GET http://localhost:8080/{location}
