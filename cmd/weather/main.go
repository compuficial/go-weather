package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"weather-cli/pkg/display"
	"weather-cli/pkg/weather"

	"github.com/fatih/color"
)

const (
	baseURL = "https://api.open-meteo.com/v1/forecast"
)

type Config struct {
	Latitude        float64
	Longitude       float64
	TemperatureUnit string
	WindSpeedUnit   string
	Timezone        string
}

var defaultConfig = Config{
	Latitude:        40.7143,
	Longitude:       -74.006,
	TemperatureUnit: "fahrenheit",
	WindSpeedUnit:   "mph",
	Timezone:        "America/New_York",
}

func buildWeatherURL(config Config) string {
	return fmt.Sprintf("%s?latitude=%f&longitude=%f&"+
		"current=temperature_2m,relative_humidity_2m,precipitation,wind_speed_10m,wind_direction_10m,wind_gusts_10m&"+
		"hourly=temperature_2m,relative_humidity_2m,precipitation_probability,wind_speed_10m&"+
		"daily=temperature_2m_max,temperature_2m_min,sunrise,sunset,precipitation_sum,precipitation_probability_max,wind_speed_10m_max&"+
		"temperature_unit=%s&"+
		"wind_speed_unit=%s&"+
		"precipitation_unit=inch&"+
		"timezone=%s",
		baseURL,
		config.Latitude,
		config.Longitude,
		config.TemperatureUnit,
		config.WindSpeedUnit,
		config.Timezone)
}

func getWeather(config Config) (*weather.WeatherResponse, error) {
	url := buildWeatherURL(config)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var weatherData weather.WeatherResponse
	if err := json.Unmarshal(body, &weatherData); err != nil {
		return nil, fmt.Errorf("failed to parse weather data: %w", err)
	}

	return &weatherData, nil
}

func main() {
	config := defaultConfig

	flag.Float64Var(&config.Latitude, "lat", defaultConfig.Latitude, "Latitude coordinate")
	flag.Float64Var(&config.Longitude, "lon", defaultConfig.Longitude, "Longitude coordinate")
	flag.StringVar(&config.TemperatureUnit, "temp", defaultConfig.TemperatureUnit, "Temperature unit (celsius, fahrenheit)")
	flag.StringVar(&config.WindSpeedUnit, "wind", defaultConfig.WindSpeedUnit, "Wind speed unit (kmh, mph, ms)")
	flag.StringVar(&config.Timezone, "tz", defaultConfig.Timezone, "Timezone (e.g., America/New_York)")
	help := flag.Bool("help", false, "Show help message")
	flag.BoolVar(help, "h", false, "Show help message (shorthand)")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	weatherData, err := getWeather(config)
	if err != nil {
		color.Red("Error: %v\n", err)
		os.Exit(1)
	}

	display.PrintCurrentWeather(weatherData)
	display.PrintHourlyForecast(weatherData)
	display.PrintDailyForecast(weatherData)
}
