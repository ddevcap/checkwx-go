package checkwx

type Barometer struct {
	Hg  float64 `json:"hg,omitempty"`
	Kpa float64 `json:"kpa,omitempty"`
	Mb  float64 `json:"mb,omitempty"`
}
