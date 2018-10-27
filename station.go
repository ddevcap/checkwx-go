package checkwx

type Station struct {
	Icao              string             `json:"icao"`
	Name              string             `json:"name"`
	Activated         string             `json:"activated,omitempty"`
	City              string             `json:"city,omitempty"`
	Country           *Country           `json:"country,omitempty"`
	Elevation         *Elevation         `json:"elevation"`
	Iata              string             `json:"iata,omitempty"`
	Latitude          *Position          `json:"latitude"`
	Longitude         *Position          `json:"longitude"`
	MagneticVariation *MagneticVariation `json:"magnetic_variation,omitempty"`
	Sectional         string             `json:"sectional,omitempty"`
	State             *State             `json:"state,omitempty"`
	Status            string             `json:"status,omitempty"`
	TimeZone          *TimeZone          `json:"timezone,omitempty"`
	Type              string             `json:"type,omitempty"`
	Usage             string             `json:"useage,omitempty"`
}
