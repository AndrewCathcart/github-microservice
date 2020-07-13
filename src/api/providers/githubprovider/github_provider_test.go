package githubprovider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthHeader(t *testing.T) {
	header := getAuthHeader("123xyz")

	assert.EqualValues(t, "token 123xyz", header)
}

func TestDefer(t *testing.T) {
	defer t.Logf("1")
	defer t.Logf("2")
	defer t.Logf("3")

	t.Logf("function's body")
}
