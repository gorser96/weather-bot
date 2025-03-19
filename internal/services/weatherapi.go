package services

import (
	"encoding/json"
	"fmt"
	"time"
	"weather-bot/internal/utils"

	"github.com/go-resty/resty/v2"
)

type WeatherapiCurrentResponse struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		WindchillC float64 `json:"windchill_c"`
		WindchillF float64 `json:"windchill_f"`
		HeatindexC float64 `json:"heatindex_c"`
		HeatindexF float64 `json:"heatindex_f"`
		DewpointC  float64 `json:"dewpoint_c"`
		DewpointF  float64 `json:"dewpoint_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
}

type hourResponse struct {
	TimeEpoch int     `json:"time_epoch"`
	Time      string  `json:"time"`
	TempC     float64 `json:"temp_c"`
	TempF     float64 `json:"temp_f"`
	IsDay     int     `json:"is_day"`
	Condition struct {
		Text string `json:"text"`
		Icon string `json:"icon"`
		Code int    `json:"code"`
	} `json:"condition"`
	WindMph      float64 `json:"wind_mph"`
	WindKph      float64 `json:"wind_kph"`
	WindDegree   int     `json:"wind_degree"`
	WindDir      string  `json:"wind_dir"`
	PressureMb   float64 `json:"pressure_mb"`
	PressureIn   float64 `json:"pressure_in"`
	PrecipMm     float64 `json:"precip_mm"`
	PrecipIn     float64 `json:"precip_in"`
	SnowCm       float64 `json:"snow_cm"`
	Humidity     int     `json:"humidity"`
	Cloud        int     `json:"cloud"`
	FeelslikeC   float64 `json:"feelslike_c"`
	FeelslikeF   float64 `json:"feelslike_f"`
	WindchillC   float64 `json:"windchill_c"`
	WindchillF   float64 `json:"windchill_f"`
	HeatindexC   float64 `json:"heatindex_c"`
	HeatindexF   float64 `json:"heatindex_f"`
	DewpointC    float64 `json:"dewpoint_c"`
	DewpointF    float64 `json:"dewpoint_f"`
	WillItRain   int     `json:"will_it_rain"`
	ChanceOfRain int     `json:"chance_of_rain"`
	WillItSnow   int     `json:"will_it_snow"`
	ChanceOfSnow int     `json:"chance_of_snow"`
	VisKm        float64 `json:"vis_km"`
	VisMiles     float64 `json:"vis_miles"`
	GustMph      float64 `json:"gust_mph"`
	GustKph      float64 `json:"gust_kph"`
	Uv           float64 `json:"uv"`
}

type WeatherapiForecastResponse struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		WindchillC float64 `json:"windchill_c"`
		WindchillF float64 `json:"windchill_f"`
		HeatindexC float64 `json:"heatindex_c"`
		HeatindexF float64 `json:"heatindex_f"`
		DewpointC  float64 `json:"dewpoint_c"`
		DewpointF  float64 `json:"dewpoint_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Date      string `json:"date"`
			DateEpoch int    `json:"date_epoch"`
			Day       struct {
				MaxtempC          float64 `json:"maxtemp_c"`
				MaxtempF          float64 `json:"maxtemp_f"`
				MintempC          float64 `json:"mintemp_c"`
				MintempF          float64 `json:"mintemp_f"`
				AvgtempC          float64 `json:"avgtemp_c"`
				AvgtempF          float64 `json:"avgtemp_f"`
				MaxwindMph        float64 `json:"maxwind_mph"`
				MaxwindKph        float64 `json:"maxwind_kph"`
				TotalprecipMm     float64 `json:"totalprecip_mm"`
				TotalprecipIn     float64 `json:"totalprecip_in"`
				TotalsnowCm       float64 `json:"totalsnow_cm"`
				AvgvisKm          float64 `json:"avgvis_km"`
				AvgvisMiles       float64 `json:"avgvis_miles"`
				Avghumidity       int     `json:"avghumidity"`
				DailyWillItRain   int     `json:"daily_will_it_rain"`
				DailyChanceOfRain int     `json:"daily_chance_of_rain"`
				DailyWillItSnow   int     `json:"daily_will_it_snow"`
				DailyChanceOfSnow int     `json:"daily_chance_of_snow"`
				Condition         struct {
					Text string `json:"text"`
					Icon string `json:"icon"`
					Code int    `json:"code"`
				} `json:"condition"`
				Uv float64 `json:"uv"`
			} `json:"day"`
			Astro struct {
				Sunrise          string `json:"sunrise"`
				Sunset           string `json:"sunset"`
				Moonrise         string `json:"moonrise"`
				Moonset          string `json:"moonset"`
				MoonPhase        string `json:"moon_phase"`
				MoonIllumination int    `json:"moon_illumination"`
				IsMoonUp         int    `json:"is_moon_up"`
				IsSunUp          int    `json:"is_sun_up"`
			} `json:"astro"`
			Hour []hourResponse `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func GetCurrentWeatherFromWeatherapi(city string, apiKey string) (string, error) {
	client := resty.New()
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"q":   city,
			"aqi": "no",
			"key": apiKey,
		}).
		Get("http://api.weatherapi.com/v1/current.json")

	if err != nil {
		return "", fmt.Errorf("failed to fetch data from Weatherapi: %v", err)
	}

	var weather WeatherapiCurrentResponse
	if err := json.Unmarshal(resp.Body(), &weather); err != nil {
		return "", fmt.Errorf("failed to parse response by Weatherapi: %v", err)
	}

	return fmt.Sprintf(
		CurrentWeatherResultStr, "Weatherapi", weather.Location.Name,
		weather.Current.TempC, weather.Current.WindMph, weather.Current.FeelslikeC), nil
}

func GetDailyWeatherFromWeatherapi(city string, apiKey string) (string, error) {
	client := resty.New()
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"q":      city,
			"aqi":    "no",
			"alerts": "no",
			"key":    apiKey,
		}).
		Get("http://api.weatherapi.com/v1/forecast.json")

	if err != nil {
		return "", fmt.Errorf("failed to fetch data by Weatherapi: %v", err)
	}

	var weather WeatherapiForecastResponse
	if err := json.Unmarshal(resp.Body(), &weather); err != nil {
		return "", fmt.Errorf("failed to parse response by Weatherapi: %v", err)
	}

	if len(weather.Forecast.Forecastday) == 0 {
		return "", fmt.Errorf("no forecast data by Weatherapi")
	}

	forecast := weather.Forecast.Forecastday[0]
	hours12 := utils.Filter(forecast.Hour, filter12ByHour)
	hours16 := utils.Filter(forecast.Hour, filter16ByHour)
	hours20 := utils.Filter(forecast.Hour, filter20ByHour)

	tempC12 := utils.MapTemps(hours12, func(item hourResponse) float64 { return item.TempC })
	tempWind12 := utils.MapTemps(hours12, func(item hourResponse) float64 { return item.WindMph })
	tempFeels12 := utils.MapTemps(hours12, func(item hourResponse) float64 { return item.FeelslikeC })
	tempC16 := utils.MapTemps(hours16, func(item hourResponse) float64 { return item.TempC })
	tempWind16 := utils.MapTemps(hours16, func(item hourResponse) float64 { return item.WindMph })
	tempFeels16 := utils.MapTemps(hours16, func(item hourResponse) float64 { return item.FeelslikeC })
	tempC20 := utils.MapTemps(hours20, func(item hourResponse) float64 { return item.TempC })
	tempWind20 := utils.MapTemps(hours20, func(item hourResponse) float64 { return item.WindMph })
	tempFeels20 := utils.MapTemps(hours20, func(item hourResponse) float64 { return item.FeelslikeC })
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
		DailyWeatherResultStr, "Weatherapi", weather.Location.Name,
		avg12Temp, avg12Wind, avg12Feels,
		avg16Temp, avg16Wind, avg16Feels,
		avg20Temp, avg20Wind, avg20Feels), nil
}

func filter12ByHour(item hourResponse) bool {
	dt := time.Unix(int64(item.TimeEpoch), 0)
	return utils.FilterHour12(dt)
}

func filter16ByHour(item hourResponse) bool {
	dt := time.Unix(int64(item.TimeEpoch), 0)
	return utils.FilterHour16(dt)
}

func filter20ByHour(item hourResponse) bool {
	dt := time.Unix(int64(item.TimeEpoch), 0)
	return utils.FilterHour20(dt)
}
