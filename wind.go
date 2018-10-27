package checkwx

type Wind struct {
	Degrees  int `json:"degrees,omitempty"`
	SpeedKts int `json:"speed_kts,omitempty"`
	SpeedMph int `json:"speed_mph,omitempty"`
	SpeedMps int `json:"speed_mps,omitempty"`
	GustKts  int `json:"gust_kts,omitempty"`
	GustMph  int `json:"gust_mph,omitempty"`
	GustMp   int `json:"gust_mp,omitempty"`
}
