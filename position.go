package checkwx

type Position struct {
	Decimal float64 `json:"decimal,omitempty"`
	Degrees string  `json:"degrees,omitempty"`
}
