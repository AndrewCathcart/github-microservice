package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "github-microservice",
		Description: "consuming external APIs in Go",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	bytes, err := json.Marshal(request)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	// Attempt to unmarshal request byte array into target
	var target CreateRepoRequest
	err = json.Unmarshal(bytes, &target)

	assert.Nil(t, err, "it should successfully unmarshal")
	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.Description, request.Description)
	assert.EqualValues(t, target.Homepage, request.Homepage)
	assert.EqualValues(t, target.Private, request.Private)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)
	assert.EqualValues(t, target.HasProjects, request.HasProjects)
	assert.EqualValues(t, target.HasWiki, request.HasWiki)
}

func TestCreateRepoResponseAsJson(t *testing.T) {
	owner := RepoOwner{
		ID:      999,
		Login:   "testingacc",
		URL:     "https://api.github.com/users/testingacc",
		HTMLURL: "https://github.com/testingacc",
	}

	permissions := RepoPermissions{
		IsAdmin: true,
		HasPull: true,
		HasPush: true,
	}

	request := CreateRepoResponse{
		ID:          123123,
		Name:        "go-test-repo",
		FullName:    "testingacc/go-test-repo",
		Owner:       owner,
		Permissions: permissions,
	}

	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	// Attempt to unmarshal request byte array into target
	var target CreateRepoResponse
	err = json.Unmarshal(bytes, &target)

	assert.Nil(t, err, "it should successfully unmarshal")
	assert.EqualValues(t, target.ID, request.ID)
	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.FullName, request.FullName)
	assert.EqualValues(t, target.Owner, request.Owner)
	assert.EqualValues(t, target.Permissions, request.Permissions)
}
