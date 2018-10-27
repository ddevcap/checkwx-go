package checkwx

type Timestamp struct {
	Issued    string `json:"issued,omitempty"`
	Bulletin  string `json:"bulletin,omitempty"`
	ValidFrom string `json:"valid_from,omitempty"`
	ValidTo   string `json:"valid_to,omitempty"`
}
