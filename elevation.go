package checkwx

type Elevation struct {
	Feet   float64 `json:"feet,omitempty"`
	Meters float64 `json:"meters,omitempty"`
	Method string  `json:"method,omitempty"`
}
