package checkwx

type Station struct {
	Icao                  string    `json:"icao"`
	Name                  string    `json:"name"`
	Activated             string    `json:"activated,omitempty"`
	City                  string    `json:"city,omitempty"`
	Country               Country   `json:"country,omitempty"`
	Elevation             Elevation `json:"elevation"`
	Iata                  string    `json:"iata,omitempty"`
	Latitude              Position  `json:"latitude"`
	Longitude             Position  `json:"longitude"`
	MagneticVariation     string    `json:"magnetic_variation,omitempty"`
	MagneticVariationYear string    `json:"magnetic_variation_year,omitempty"`
	Sectional             string    `json:"sectional,omitempty"`
	State                 string    `json:"state,omitempty"`
	Status                string    `json:"status,omitempty"`
	TimeZone              TimeZone  `json:"time_zone"`
	Type                  string    `json:"type,omitempty"`
	Usage                 string    `json:"usage,omitempty"`
}
