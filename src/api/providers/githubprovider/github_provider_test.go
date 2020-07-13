package githubprovider

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/andrewcathcart/github-microservice/src/api/clients/restclient"
	"github.com/andrewcathcart/github-microservice/src/api/domain/github"
	"github.com/stretchr/testify/assert"
)

var (
	restClientMock restClientMockStruct
	postMock       func(url string, body interface{}, headers http.Header) (*http.Response, error)
)

type restClientMockStruct struct{}

func (m *restClientMockStruct) Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	return postMock(url, body, headers)
}

func init() {
	restclient.RestClient = &restClientMockStruct{}
}

func TestGetAuthHeader(t *testing.T) {
	header := getAuthHeader("123xyz")

	assert.EqualValues(t, "token 123xyz", header)
}

func TestCreateRepo(t *testing.T) {
	postMock = func(url string, body interface{}, headers http.Header) (*http.Response, error) {
		return nil, fmt.Errorf("Auth missing")
	}

	response, err := CreateRepo("", &github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
}
