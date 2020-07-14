package githubprovider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/andrewcathcart/github-microservice/src/api/clients/restclient"
	"github.com/andrewcathcart/github-microservice/src/api/domain/github"
)

const (
	headerAuth       = "Authorization"
	headerAuthFormat = "token %s"
	urlCreateRepo    = "https://api.github.com/user/repos"
)

func getAuthHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthFormat, accessToken)
}

// CreateRepo sends a POST request to the github api to create a new repository
func CreateRepo(accessToken string, req *github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuth, getAuthHeader(accessToken))

	response, err := restclient.RestClient.Post(urlCreateRepo, req, headers)
	fmt.Println(response)
	if err != nil {
		log.Println(fmt.Sprintf("An error occurred when trying to create new repo in github: %s", err.Error()))
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "Invalid response body."}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "Invalid json response body."}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("An error occurred when trying to unmarshal CreateRepo successful response: %s", err.Error()))
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "An error occurred when trying to unmarshal github CreateRepo response."}
	}

	return &result, nil
}
