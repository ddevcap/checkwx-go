package checkwx

import "encoding/json"

type ApiResponse struct {
	Results int             `json:"results"`
	Data    interface{}     `json:"data"`
	Errors  []ErrorResponse `json:"errors"`
}

func (ar *ApiResponse) ToStation() ([]*Station, error) {
	var s []*Station
	err := ar.As(&s)
	return s, err
}

func (ar *ApiResponse) ToMetar() ([]*Metar, error) {
	var m []*Metar
	err := ar.As(&m)
	return m, err
}

func (ar *ApiResponse) ToTaf() ([]*Taf, error) {
	var t []*Taf
	err := ar.As(&t)
	return t, err
}

func (ar *ApiResponse) As(v interface{}) error {
	b, err := json.Marshal(ar.Data)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &v)
}
