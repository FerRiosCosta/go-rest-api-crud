//go:build e2e
// +build e2e

package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetComments(t *testing.T) {
	fmt.Println("Running E2E test fot get comments endpoint")

	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/comment")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}

func TestPostComments(t *testing.T) {
	fmt.Println("Running E2E test fot POST comments endpoint")

	client := resty.New()
	resp, err := client.R().
		SetBody(`{"slug": "/test", "author": "test", "body": "hello world"}`).
		Post(BASE_URL + "/api/comment")

	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
}
