package checkwx

type DewPoint struct {
	Celsius    int `json:"celsius,omitempty"`
	Fahrenheit int `json:"fahrenheit,omitempty"`
}
