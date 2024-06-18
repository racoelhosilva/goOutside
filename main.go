package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Weather struct {
	Location struct {
		City    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Forecast struct {
		ForecastDay []struct {
			Day struct {
				Tides []struct {
					Tide []struct {
						Time string `json:"tide_time"`
						Type string `json:"tide_type"`
					} `json:"tide"`
				} `json:"tides"`
			} `json:"day"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	city := "Oporto"

	resource := fmt.Sprintf("http://api.weatherapi.com/v1/marine.json?key=%s&q=%s&days=1", os.Getenv("WEATHER_API_KEY"), city)
	response, err := http.Get(resource)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	var weather Weather
	json.Unmarshal([]byte(body), &weather)

	fmt.Println(weather.Location.City, weather.Location.Country)
	for _, tide := range weather.Forecast.ForecastDay[0].Day.Tides[0].Tide {
		fmt.Printf("%-10s %s\n", tide.Type, tide.Time)
	}
}
