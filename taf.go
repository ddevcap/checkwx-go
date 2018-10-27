package checkwx

type Taf struct {
	Icao      string     `json:"icao"`
	Timestamp Timestamp  `json:"timestamp"`
	RawText   string     `json:"raw_text"`
	Forecast  []Forecast `json:"forecast,omitempty"`
}
