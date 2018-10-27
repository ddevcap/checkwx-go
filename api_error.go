package checkwx

import "net/http"

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Help    string `json:"help"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

type ApiError struct {
	req *http.Request
	res *http.Response
	err *ErrorResponse
}

type ApiRateLimitError struct {
	ApiError
}

func (e ApiRateLimitError) Error() string {
	return e.ApiError.err.Message
}

type InvalidHeaderApiKeyError struct {
	ApiError
}

func (e InvalidHeaderApiKeyError) Error() string {
	return e.ApiError.err.Message
}

type ServerError struct {
	ApiError
}

func (e ServerError) Error() string {
	return e.res.Status
}

type UnknownError struct {
	ApiError
}

func (e UnknownError) Error() string {
	return e.res.Status
}
