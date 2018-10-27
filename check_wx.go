package checkwx

import (
	"encoding/json"
	"fmt"
	"github.com/moul/http2curl"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type CheckWx struct {
	client      *http.Client
	apiKey      string
	Debug       bool
	QueryParams map[string]string
	Headers     map[string]string
	BaseURL     string

	Station *StationService
	Metar   *MetarService
	Taf     *TafService
}

func NewCheckWx(apiKey string) *CheckWx {
	cw := &CheckWx{
		client: http.DefaultClient,
		apiKey: apiKey,
		Debug:  false,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"X-API-Key":    apiKey,
		},
		BaseURL: "https://api.checkwx.com",
	}

	cw.Station = &StationService{cw: cw}
	cw.Metar = &MetarService{cw: cw}
	cw.Taf = &TafService{cw: cw}

	return cw
}

func (c *CheckWx) newRequest(method, path string, query url.Values, body io.Reader) (*http.Request, error) {
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}

	for key, value := range c.QueryParams {
		query.Set(key, value)
	}

	u.Path = path
	u.RawQuery = query.Encode()

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

func (c *CheckWx) do(req *http.Request, v interface{}) error {
	if c.Debug == true {
		command, _ := http2curl.GetCurlCommand(req)
		fmt.Println(command)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode >= 200 && res.StatusCode < 400 {
		if v != nil {
			defer res.Body.Close()
			err = json.NewDecoder(res.Body).Decode(&v)
			if err != nil {
				return err
			}
		}

		return nil
	}

	apiError := c.handleError(req, res)

	if apiError != nil {
		return apiError
	}

	return c.do(req, v)
}

func (c *CheckWx) handleError(req *http.Request, res *http.Response) error {
	if c.Debug == true {
		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%q", dump)
	}

	var e ErrorResponse
	defer res.Body.Close()
	err := json.NewDecoder(res.Body).Decode(&e)
	if err != nil {
		return err
	}

	apiError := ApiError{
		req: req,
		res: res,
		err: &e,
	}

	switch status := apiError.res.StatusCode; status {
	case 401:
		return InvalidHeaderApiKeyError{apiError}
	case 429:
		return ApiRateLimitError{apiError}
	case 500:
		return ServerError{apiError}
	default:
		return UnknownError{apiError}
	}

	return e
}
