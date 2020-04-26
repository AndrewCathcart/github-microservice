package github_provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthHeader(t *testing.T) {
	header := getAuthHeader("123xyz")

	assert.EqualValues(t, "token 123xyz", header)
}
