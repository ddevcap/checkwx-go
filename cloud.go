package checkwx

type Cloud struct {
	Code          string  `json:"code,omitempty"`
	Text          string  `json:"text,omitempty"`
	BaseFeetAgl   float64 `json:"base_feet_agl,omitempty"`
	BaseMetersAgl float64 `json:"base_meters_agl,omitempty"`
}
