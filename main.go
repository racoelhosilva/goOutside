package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type MarineWeather struct {
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
	Forecast struct {
		Forecastday []struct {
			Date      string `json:"date"`
			DateEpoch int    `json:"date_epoch"`
			Day       struct {
				MaxtempC      float64 `json:"maxtemp_c"`
				MaxtempF      float64 `json:"maxtemp_f"`
				MintempC      float64 `json:"mintemp_c"`
				MintempF      float64 `json:"mintemp_f"`
				AvgtempC      float64 `json:"avgtemp_c"`
				AvgtempF      float64 `json:"avgtemp_f"`
				MaxwindMph    float64 `json:"maxwind_mph"`
				MaxwindKph    float64 `json:"maxwind_kph"`
				TotalprecipMm float64 `json:"totalprecip_mm"`
				TotalprecipIn float64 `json:"totalprecip_in"`
				TotalsnowCm   float64 `json:"totalsnow_cm"`
				AvgvisKm      float64 `json:"avgvis_km"`
				AvgvisMiles   float64 `json:"avgvis_miles"`
				Avghumidity   float64 `json:"avghumidity"`
				Tides         []struct {
					Tide []struct {
						TideTime     string `json:"tide_time"`
						TideHeightMt string `json:"tide_height_mt"`
						TideType     string `json:"tide_type"`
					} `json:"tide"`
				} `json:"tides"`
				Condition struct {
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
			Hour []struct {
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
				WindMph         float64 `json:"wind_mph"`
				WindKph         float64 `json:"wind_kph"`
				WindDegree      int     `json:"wind_degree"`
				WindDir         string  `json:"wind_dir"`
				PressureMb      float64 `json:"pressure_mb"`
				PressureIn      float64 `json:"pressure_in"`
				PrecipMm        float64 `json:"precip_mm"`
				PrecipIn        float64 `json:"precip_in"`
				Humidity        int     `json:"humidity"`
				Cloud           int     `json:"cloud"`
				FeelslikeC      float64 `json:"feelslike_c"`
				FeelslikeF      float64 `json:"feelslike_f"`
				WindchillC      float64 `json:"windchill_c"`
				WindchillF      float64 `json:"windchill_f"`
				HeatindexC      float64 `json:"heatindex_c"`
				HeatindexF      float64 `json:"heatindex_f"`
				DewpointC       float64 `json:"dewpoint_c"`
				DewpointF       float64 `json:"dewpoint_f"`
				VisKm           float64 `json:"vis_km"`
				VisMiles        float64 `json:"vis_miles"`
				GustMph         float64 `json:"gust_mph"`
				GustKph         float64 `json:"gust_kph"`
				Uv              float64 `json:"uv"`
				SigHtMt         float64 `json:"sig_ht_mt"`
				SwellHtMt       float64 `json:"swell_ht_mt"`
				SwellHtFt       float64 `json:"swell_ht_ft"`
				SwellDir        float64 `json:"swell_dir"`
				SwellDir16Point string  `json:"swell_dir_16_point"`
				SwellPeriodSecs float64 `json:"swell_period_secs"`
				WaterTempC      float64 `json:"water_temp_c"`
				WaterTempF      float64 `json:"water_temp_f"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	city := "Oporto"

	resource := fmt.Sprintf("http://api.weatherapi.com/v1/marine.json?key=%s&q=%s&days=7", os.Getenv("WEATHER_API_KEY"), city)
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

	var weather MarineWeather
	json.Unmarshal([]byte(body), &weather)

	fmt.Println(weather.Location.Name, weather.Location.Country)
	for _, day := range weather.Forecast.Forecastday {
		fmt.Printf("%-12s %.1fºC  %s\n", day.Date, day.Day.AvgtempC, day.Day.Condition.Text)
		
		astroLayout := "03:04 PM"
		sunrise, err := time.Parse(astroLayout, day.Astro.Sunrise)
		if err != nil {
			panic(err)
		}
		sunset, err := time.Parse(astroLayout, day.Astro.Sunset)
		if err != nil {
			panic(err)
		}

		fmt.Printf("  Sunrise: %02d:%02d  Sunset: %02d:%02d\n", sunrise.Hour(), sunrise.Minute(), sunset.Hour(), sunset.Minute())
		
		for _, tide := range day.Day.Tides[0].Tide {
			layout := "2006-01-02 15:04"
			t, err := time.Parse(layout, tide.TideTime)
		
			if err != nil {
				panic(err)
			}
		
			fmt.Printf("    %-10s %02d:%02d\n", tide.TideType, t.Hour(), t.Minute())
		}
	}
}
