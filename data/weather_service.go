package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	weathernext "github.com/tomharris/weather-next"
)

type WeatherService struct {
	ApiKey      string
	ApiEndpoint string
}

type apiResponse struct {
	Latitude  float64            `json:"lat"`
	Longitude float64            `json:"lon"`
	Alerts    []apiResponseAlert `json:"alerts"`
	Current   struct {
		Temperature float64 `json:"temp"`
		FeelsLike   float64 `json:"feels_like"`
		Weather     []struct {
			Main        string `json:"main"`
			Description string `json:"description"`
		} `json:"weather"`
	} `json:"current"`
}

type apiResponseAlert struct {
	Event       string `json:"event"`
	Description string `json:"description"`
	Sender      string `json:"sender_name"`
	Start       int64  `json:"start"`
	End         int64  `json:"end"`
}

func (ws WeatherService) GetWeatherForLocation(latitude float64, longitude float64) (weathernext.LocationWeather, error) {
	params := url.Values{}
	params.Add("appid", ws.ApiKey)
	params.Add("lat", fmt.Sprintf("%v", latitude))
	params.Add("lon", fmt.Sprintf("%v", longitude))
	params.Add("units", "imperial")
	params.Add("exclude", "minutely,hourly,daily")

	resp, err := http.Get(ws.ApiEndpoint + "?" + params.Encode())
	if err != nil {
		return weathernext.LocationWeather{}, fmt.Errorf("error contacting weather api: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return weathernext.LocationWeather{}, fmt.Errorf("error reading response body: %v", err)
	}

	return extractToLocationWeather(responseBody), nil
}

func extractToLocationWeather(responseJson []byte) weathernext.LocationWeather {
	var parsedJson apiResponse
	json.Unmarshal(responseJson, &parsedJson)

	return weathernext.LocationWeather{
		Latitude:           parsedJson.Latitude,
		Longitude:          parsedJson.Longitude,
		Temperature:        parsedJson.Current.Temperature,
		FeelsLike:          parsedJson.Current.FeelsLike,
		WeatherType:        firstWeatherType(parsedJson),
		WeatherDescription: firstWeatherDescription(parsedJson),
		Alerts:             extractToWeatherAlerts(parsedJson.Alerts),
	}
}

func extractToWeatherAlerts(responseAlerts []apiResponseAlert) []weathernext.WeatherAlert {
	var alerts []weathernext.WeatherAlert
	for _, alert := range responseAlerts {
		alerts = append(alerts, weathernext.WeatherAlert{Sender: alert.Sender, Event: alert.Event, Start: alert.Start, End: alert.End, Description: alert.Description})
	}
	return alerts
}

func firstWeatherType(parsedJson apiResponse) string {
	if len(parsedJson.Current.Weather) == 0 {
		return ""
	}
	return parsedJson.Current.Weather[0].Main
}

func firstWeatherDescription(parsedJson apiResponse) string {
	if len(parsedJson.Current.Weather) == 0 {
		return ""
	}
	return parsedJson.Current.Weather[0].Description
}
