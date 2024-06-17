package api

import (
	"fmt"
	"io"
	"os"
	"net/http"
)

func GetCurrent(city string) string {
	resource := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", os.Getenv("WEATHER_API_KEY"), city)
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

	return string(body)
}
