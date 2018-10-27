package checkwx

import (
	"fmt"
	"net/http"
	"strings"
)

type TafService Service

func (service *TafService) GetTaf(icaos string) (*Taf, error) {
	metars, err := service.GetTafs(icaos)
	if err != nil {
		return nil, err
	}
	return metars[0], err
}

func (service *TafService) GetTafs(icaos ...string) ([]*Taf, error) {
	path := fmt.Sprintf("/taf/%s/decoded", strings.Join(icaos, ","))
	method := http.MethodGet

	req, err := service.cw.newRequest(method, path, nil, nil)
	if err != nil {
		return nil, err
	}

	var res ApiResponse
	if err := service.cw.do(req, &res); err != nil {
		return nil, err
	}

	return res.ToTaf()
}
