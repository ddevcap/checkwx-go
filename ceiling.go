package checkwx

type Ceiling struct {
	Code      string  `json:"code,omitempty"`
	Text      string  `json:"text,omitempty"`
	FeetAgl   float64 `json:"feet_agl,omitempty"`
	MetersAgl float64 `json:"meters_agl,omitempty"`
}
