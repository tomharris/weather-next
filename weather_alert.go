package weathernext

type WeatherAlert struct {
	Sender      string `json:"sender"`
	Event       string `json:"event"`
	Description string `json:"description"`
	Start       int64  `json:"start"`
	End         int64  `json:"end"`
}
