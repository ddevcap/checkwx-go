package checkwx

type Metar struct {
	Icao            string      `json:"icao"`
	Name            string      `json:"name"`
	Observed        string      `json:"observed"`
	RawText         string      `json:"raw_text"`
	Barometer       Barometer   `json:"barometer,omitempty"`
	Ceiling         Ceiling     `json:"ceiling,omitempty"`
	Clouds          []Cloud     `json:"clouds,omitempty"`
	Conditions      []Condition `json:"conditions,omitempty"`
	DewPoint        DewPoint    `json:"dewpoint,omitempty"`
	Elevation       Elevation   `json:"elevation,omitempty"`
	FlightCategory  string      `json:"flight_category,omitempty"`
	HumidityPercent int         `json:"humidity_percent,omitempty"`
	RainIn          int         `json:"rain_in,omitempty"`
	SnowIn          int         `json:"snow_in,omitempty"`
	Temperature     Temperature `json:"temperature,omitempty"`
	Visibility      Visibility  `json:"visibility,omitempty"`
	Wind            Wind        `json:"wind,omitempty"`
}
