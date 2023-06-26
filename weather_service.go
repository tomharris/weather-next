package weathernext

type LocationWeatherService interface {
	GetWeatherForLocation(latitude float64, longitude float64) (LocationWeather, error)
}
