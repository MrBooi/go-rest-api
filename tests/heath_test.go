//go:build end_two_end
// +build end_two_end

package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	fmt.Println("Running end two end test for health check endpoint")

	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/health")
	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
}
