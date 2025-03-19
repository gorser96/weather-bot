package services

const DailyWeatherResultStr = "(%s) Temperature in %s:\n" +
	"\xF0\x9F\x95\x92 08:00 - 12:00 -> %.1f°C | \xf0\x9f\x92\xa8 %.1f м/с | \U0001f624 %.1f°C \n" +
	"\xF0\x9F\x95\x97 12:00 - 16:00 -> %.1f°C | \xf0\x9f\x92\xa8 %.1f м/с | \U0001f624 %.1f°C \n" +
	"\xF0\x9F\x95\x99 16:00 - 21:00 -> %.1f°C | \xf0\x9f\x92\xa8 %.1f м/с | \U0001f624 %.1f°C"
const CurrentWeatherResultStr = "(%s) Temperature in %s: %.1f°C | \xf0\x9f\x92\xa8 %.1f м/с | \U0001f624 %.1f°C"

type WeatherService interface {
	GetCurrentWeather(city string) (string, error)
	GetDailyWeather(city string) (string, error)
}

type OpenWeatherService struct {
	APIKey string
}

type WeatherapiService struct {
	APIKey string
}

type OpenmeteoService struct {
	OpenWeatherApiKey string
}

func (s OpenWeatherService) GetCurrentWeather(city string) (string, error) {
	return GetCurrentWeatherFromOpenWeather(city, s.APIKey)
}

func (s OpenWeatherService) GetDailyWeather(city string) (string, error) {
	return GetDailyWeatherFromOpenWeather(city, s.APIKey)
}

func (s WeatherapiService) GetCurrentWeather(city string) (string, error) {
	return GetCurrentWeatherFromWeatherapi(city, s.APIKey)
}

func (s WeatherapiService) GetDailyWeather(city string) (string, error) {
	return GetDailyWeatherFromWeatherapi(city, s.APIKey)
}

func (s OpenmeteoService) GetCurrentWeather(city string) (string, error) {
	return GetCurrentWeatherFromOpenmeteo(city, s.OpenWeatherApiKey)
}

func (s OpenmeteoService) GetDailyWeather(city string) (string, error) {
	return GetDailyWeatherFromOpenmeteo(city, s.OpenWeatherApiKey)
}
