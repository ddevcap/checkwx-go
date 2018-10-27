package checkwx

type TimeZone struct {
	Gmt  int    `json:"gmt,omitempty"`
	Dst  int    `json:"dst,omitempty"`
	Tzid string `json:"tzid,omitempty"`
}
