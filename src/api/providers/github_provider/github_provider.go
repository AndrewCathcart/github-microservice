package github_provider

import (
	"fmt"

	"github.com/andrewcathcart/github-microservice/src/api/domain/github"
)

const (
	headerAuth       = "Authorization"
	headerAuthFormat = "token %s"
)

func getAuthHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthFormat, accessToken)
}

func CreateRepo(accessToken string, req *github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	header := getAuthHeader(accessToken)
	return &github.CreateRepoResponse{}, &github.GithubErrorResponse{}
}
