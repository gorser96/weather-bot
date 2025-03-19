package services

import (
	"fmt"
	"log"
	"time"
	"weather-bot/internal/utils"

	owm "github.com/briandowns/openweathermap"
)

func GetCurrentWeatherFromOpenWeather(city string, apiKey string) (string, error) {
	weather, err := owm.NewCurrent("C", "ru", apiKey)
	if err != nil {
		log.Fatalln(err)
		return "", fmt.Errorf("failed to connect with OpenWeatherMap client: %v", err)
	}

	err = weather.CurrentByName(city)
	if err != nil {
		log.Fatalln(err)
		return "", fmt.Errorf("failed to get current weather by OpenWeatherMap: %v", err)
	}

	return fmt.Sprintf(
		CurrentWeatherResultStr, "OpenWeatherMap", weather.Name,
		weather.Main.Temp, weather.Wind.Speed, weather.Main.FeelsLike), nil
}

func GetDailyWeatherFromOpenWeather(city string, apiKey string) (string, error) {
	weather, err := owm.NewForecast("5", "C", "ru", apiKey)
	if err != nil {
		log.Fatalln(err)
		return "", fmt.Errorf("failed to connect with OpenWeatherMap client: %v", err)
	}

	err = weather.DailyByName(city, 8)
	if err != nil {
		log.Fatalln(err)
		return "", fmt.Errorf("failed to get daily weather by OpenWeatherMap: %v", err)
	}

	forecastResponse := weather.ForecastWeatherJson.(*owm.Forecast5WeatherData)
	if forecastResponse == nil {
		return "", fmt.Errorf("failed to unpacking daily weather by OpenWeatherMap")
	}

	if len(forecastResponse.List) == 0 {
		return "", fmt.Errorf("no forecast data by OpenWeatherMap")
	}

	hours12 := utils.Filter(forecastResponse.List, filter12ByOwm)
	hours16 := utils.Filter(forecastResponse.List, filter16ByOwm)
	hours20 := utils.Filter(forecastResponse.List, filter20ByOwm)

	tempC12 := utils.MapTemps(hours12, func(item owm.Forecast5WeatherList) float64 { return item.Main.Temp })
	tempWind12 := utils.MapTemps(hours12, func(item owm.Forecast5WeatherList) float64 { return item.Wind.Speed })
	tempFeels12 := utils.MapTemps(hours12, func(item owm.Forecast5WeatherList) float64 { return item.Main.FeelsLike })
	tempC16 := utils.MapTemps(hours16, func(item owm.Forecast5WeatherList) float64 { return item.Main.Temp })
	tempWind16 := utils.MapTemps(hours16, func(item owm.Forecast5WeatherList) float64 { return item.Wind.Speed })
	tempFeels16 := utils.MapTemps(hours16, func(item owm.Forecast5WeatherList) float64 { return item.Main.FeelsLike })
	tempC20 := utils.MapTemps(hours20, func(item owm.Forecast5WeatherList) float64 { return item.Main.Temp })
	tempWind20 := utils.MapTemps(hours20, func(item owm.Forecast5WeatherList) float64 { return item.Wind.Speed })
	tempFeels20 := utils.MapTemps(hours20, func(item owm.Forecast5WeatherList) float64 { return item.Main.FeelsLike })
	avg12Temp := utils.Avg(tempC12)
	avg12Wind := utils.Avg(tempWind12)
	avg12Feels := utils.Avg(tempFeels12)
	avg16Temp := utils.Avg(tempC16)
	avg16Wind := utils.Avg(tempWind16)
	avg16Feels := utils.Avg(tempFeels16)
	avg20Temp := utils.Avg(tempC20)
	avg20Wind := utils.Avg(tempWind20)
	avg20Feels := utils.Avg(tempFeels20)

	return fmt.Sprintf(
		DailyWeatherResultStr, "OpenWeatherMap", forecastResponse.City.Name,
		avg12Temp, avg12Wind, avg12Feels,
		avg16Temp, avg16Wind, avg16Feels,
		avg20Temp, avg20Wind, avg20Feels), nil
}

func filter12ByOwm(item owm.Forecast5WeatherList) bool {
	dt := time.Unix(int64(item.Dt), 0)
	return utils.FilterHour12(dt.UTC())
}

func filter16ByOwm(item owm.Forecast5WeatherList) bool {
	dt := time.Unix(int64(item.Dt), 0)
	return utils.FilterHour16(dt.UTC())
}

func filter20ByOwm(item owm.Forecast5WeatherList) bool {
	dt := time.Unix(int64(item.Dt), 0)
	return utils.FilterHour20(dt.UTC())
}
