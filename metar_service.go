package checkwx

import (
	"fmt"
	"net/http"
	"strings"
)

type MetarService Service

func (service *MetarService) GetMetar(icaos string) (*Metar, error) {
	metars, err := service.GetMetars(icaos)
	if err != nil {
		return nil, err
	}
	return metars[0], err
}

func (service *MetarService) GetMetars(icaos ...string) ([]*Metar, error) {
	path := fmt.Sprintf("/metar/%s/decoded", strings.Join(icaos, ","))
	method := http.MethodGet

	req, err := service.cw.newRequest(method, path, nil, nil)
	if err != nil {
		return nil, err
	}

	var res ApiResponse
	if err := service.cw.do(req, &res); err != nil {
		return nil, err
	}

	return res.ToMetar()
}
