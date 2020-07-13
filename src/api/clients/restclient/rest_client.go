package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var (
	RestClient restClientInterface
)

type restClientInterface interface {
	Post(url string, body interface{}, headers http.Header) (*http.Response, error)
}

type restClient struct{}

func (u *restClient) Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}

func init() {
	RestClient = &restClient{}
}
