package checkwx

type Temperature struct {
	Celsius    int `json:"celsius,omitempty"`
	Fahrenheit int `json:"fahrenheit,omitempty"`
}
