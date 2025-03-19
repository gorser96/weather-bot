package services

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type openWeatherMapGeoResponse struct {
	Name       string `json:"name"`
	LocalNames struct {
		Sr          string `json:"sr"`
		Ru          string `json:"ru"`
		Os          string `json:"os"`
		Hy          string `json:"hy"`
		Lt          string `json:"lt"`
		Et          string `json:"et"`
		Zh          string `json:"zh"`
		FeatureName string `json:"feature_name"`
		El          string `json:"el"`
		ASCII       string `json:"ascii"`
		Ka          string `json:"ka"`
		Az          string `json:"az"`
		Sl          string `json:"sl"`
		De          string `json:"de"`
		Nl          string `json:"nl"`
		En          string `json:"en"`
		Hr          string `json:"hr"`
		Uk          string `json:"uk"`
		Ro          string `json:"ro"`
		Hu          string `json:"hu"`
		Ml          string `json:"ml"`
		Sk          string `json:"sk"`
		Fr          string `json:"fr"`
		Ko          string `json:"ko"`
		Tr          string `json:"tr"`
		Ca          string `json:"ca"`
		Pl          string `json:"pl"`
	} `json:"local_names,omitempty"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
	State   string  `json:"state"`
}

func getGeoLocation(city string, apiKey string) (*openWeatherMapGeoResponse, error) {
	client := resty.New()
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"q":     city,
			"limit": "5",
			"appid": apiKey,
		}).
		Get("http://api.openweathermap.org/geo/1.0/direct")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from OpenWeatherMapGeo: %v", err)
	}

	var geoLocations []openWeatherMapGeoResponse
	if err := json.Unmarshal(resp.Body(), &geoLocations); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	if len(geoLocations) == 0 {
		return nil, fmt.Errorf("geo location for %s not found", city)
	}

	return &geoLocations[0], nil
}
