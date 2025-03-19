package services

import (
	"encoding/json"
	"fmt"
	"strconv"
	"weather-bot/internal/utils"

	"github.com/go-resty/resty/v2"
)

type OpenmeteoCurrentResponse struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	CurrentUnits         struct {
		Time                string `json:"time"`
		Interval            string `json:"interval"`
		Temperature2M       string `json:"temperature_2m"`
		WindSpeed10M        string `json:"wind_speed_10m"`
		ApparentTemperature string `json:"apparent_temperature"`
	} `json:"current_units"`
	Current struct {
		Time                string  `json:"time"`
		Interval            int     `json:"interval"`
		Temperature2M       float64 `json:"temperature_2m"`
		WindSpeed10M        float64 `json:"wind_speed_10m"`
		ApparentTemperature float64 `json:"apparent_temperature"`
	} `json:"current"`
}

type OpenmeteoDailyResponse struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	HourlyUnits          struct {
		Time                string `json:"time"`
		Temperature2M       string `json:"temperature_2m"`
		WindSpeed10M        string `json:"wind_speed_10m"`
		ApparentTemperature string `json:"apparent_temperature"`
	} `json:"hourly_units"`
	Hourly struct {
		Time                []string  `json:"time"`
		Temperature2M       []float64 `json:"temperature_2m"`
		WindSpeed10M        []float64 `json:"wind_speed_10m"`
		ApparentTemperature []float64 `json:"apparent_temperature"`
	} `json:"hourly"`
}

func GetCurrentWeatherFromOpenmeteo(city string, openWeatherApiKey string) (string, error) {
	geoLocation, err := getGeoLocation(city, openWeatherApiKey)
	if err != nil {
		return "", err
	}

	client := resty.New()
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"latitude":        strconv.FormatFloat(geoLocation.Lat, 'f', 4, 64),
			"longitude":       strconv.FormatFloat(geoLocation.Lon, 'f', 4, 64),
			"current":         "temperature_2m,wind_speed_10m,apparent_temperature",
			"forecast_days":   "1",
			"wind_speed_unit": "ms",
		}).
		Get("https://api.open-meteo.com/v1/forecast")

	if err != nil {
		return "", fmt.Errorf("failed to fetch data from Openmeteo: %v", err)
	}

	var weather OpenmeteoCurrentResponse
	if err := json.Unmarshal(resp.Body(), &weather); err != nil {
		return "", fmt.Errorf("failed to parse response by Openmeteo: %v", err)
	}

	return fmt.Sprintf(
		CurrentWeatherResultStr, "Openmeteo", geoLocation.Name,
		weather.Current.Temperature2M, weather.Current.WindSpeed10M, weather.Current.ApparentTemperature), nil
}

func GetDailyWeatherFromOpenmeteo(city string, openWeatherApiKey string) (string, error) {
	geoLocation, err := getGeoLocation(city, openWeatherApiKey)
	if err != nil {
		return "", err
	}

	client := resty.New()
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"latitude":        strconv.FormatFloat(geoLocation.Lat, 'f', 4, 64),
			"longitude":       strconv.FormatFloat(geoLocation.Lon, 'f', 4, 64),
			"hourly":          "temperature_2m,wind_speed_10m,apparent_temperature",
			"models":          "metno_seamless",
			"forecast_days":   "1",
			"wind_speed_unit": "ms",
		}).
		Get("https://api.open-meteo.com/v1/forecast")

	if err != nil {
		return "", fmt.Errorf("failed to fetch data from Openmeteo: %v", err)
	}

	var weather OpenmeteoDailyResponse
	if err := json.Unmarshal(resp.Body(), &weather); err != nil {
		return "", fmt.Errorf("failed to parse response by Openmeteo: %v", err)
	}

	avg12Temp := utils.Avg(weather.Hourly.Temperature2M[7:13])
	avg12Wind := utils.Avg(weather.Hourly.WindSpeed10M[7:13])
	avg12Feels := utils.Avg(weather.Hourly.ApparentTemperature[7:13])
	avg16Temp := utils.Avg(weather.Hourly.Temperature2M[12:17])
	avg16Wind := utils.Avg(weather.Hourly.WindSpeed10M[12:17])
	avg16Feels := utils.Avg(weather.Hourly.ApparentTemperature[12:17])
	avg20Temp := utils.Avg(weather.Hourly.Temperature2M[16:22])
	avg20Wind := utils.Avg(weather.Hourly.WindSpeed10M[16:22])
	avg20Feels := utils.Avg(weather.Hourly.ApparentTemperature[16:22])

	return fmt.Sprintf(
		DailyWeatherResultStr, "Openmeteo", geoLocation.Name,
		avg12Temp, avg12Wind, avg12Feels,
		avg16Temp, avg16Wind, avg16Feels,
		avg20Temp, avg20Wind, avg20Feels), nil
}
