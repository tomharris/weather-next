package weathernext

type LocationWeather struct {
	Latitude           float64        `json:"latitude"`
	Longitude          float64        `json:"longitude"`
	Temperature        float64        `json:"temperature"`
	FeelsLike          float64        `json:"feels_like"`
	WeatherType        string         `json:"weather_type"`
	WeatherDescription string         `json:"weather_description"`
	Alerts             []WeatherAlert `json:"alerts"`
}
