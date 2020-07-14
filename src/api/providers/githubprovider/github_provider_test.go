package githubprovider

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

func TestCreateRepoRestclientError(t *testing.T) {
	postMock = func(url string, body interface{}, headers http.Header) (*http.Response, error) {
		return nil, errors.New("Auth missing")
	}

	response, err := CreateRepo("", &github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "Auth missing", err.Message)
}

func TestCreateRepoInvalidResponseBody(t *testing.T) {
	invalidCloser, _ := os.Open("-sdfisq9w8")
	postMock = func(url string, body interface{}, headers http.Header) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusCreated, Body: invalidCloser}, nil
	}

	response, err := CreateRepo("", &github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Invalid response body.", err.Message)
}

func TestCreateRepoInvalidInterfaceError(t *testing.T) {
	postMock = func(url string, body interface{}, headers http.Header) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusUnauthorized, Body: ioutil.NopCloser(strings.NewReader(`{"message": 1}`))}, nil
	}

	response, err := CreateRepo("", &github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Invalid json response body.", err.Message)
}

func TestCreateRepoUnauthorized(t *testing.T) {
	postMock = func(url string, body interface{}, headers http.Header) (*http.Response, error) {
		return &http.Response{
				StatusCode: http.StatusUnauthorized,
				Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url": "https://developer.github.com/v3/repos/#create"}`)),
			},
			nil
	}

	response, err := CreateRepo("", &github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "Requires authentication", err.Message)
}
