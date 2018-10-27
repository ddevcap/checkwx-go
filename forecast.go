package checkwx

type Forecast struct {
	Timestamp       ForecastTimestamp `json:"timestamp,omitempty,omitempty"`
	Clouds          []Cloud           `json:"clouds,omitempty"`
	Conditions      []Condition       `json:"conditions"`
	Visibility      Visibility        `json:"visibility,omitempty"`
	Wind            Wind              `json:"wind,omitempty"`
	ChangeIndicator string            `json:"change_indicator,omitempty"`
}
