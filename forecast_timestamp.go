package checkwx

type ForecastTimestamp struct {
	ForecastFrom string `json:"forecast_from,omitempty"`
	ForecastTo   string `json:"forecast_to,omitempty"`
}
