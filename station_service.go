package checkwx

import (
	"fmt"
	"net/http"
	"strings"
)

type StationService Service

func (service *StationService) GetStation(icaos string) (*Station, error) {
	stations, err := service.GetStations(icaos)
	if err != nil {
		return nil, err
	}
	return stations[0], err
}

func (service *StationService) GetStations(icaos ...string) ([]*Station, error) {
	path := fmt.Sprintf("/station/%s/", strings.Join(icaos, ","))
	method := http.MethodGet

	req, err := service.cw.newRequest(method, path, nil, nil)
	if err != nil {
		return nil, err
	}

	var res ApiResponse
	if err := service.cw.do(req, &res); err != nil {
		return nil, err
	}

	return res.ToStation()
}
