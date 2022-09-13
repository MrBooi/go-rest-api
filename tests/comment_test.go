//go:build end_two_end
// +build end_two_end

package test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetComments(t *testing.T) {

	t.Run("Get comment by id", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().SetHeader("Content-Type", "application/json").Get(BASE_URL + "/api/v1/comment/123")
		if err != nil {
			t.Fail()
		}

		assert.Equal(t, 500, resp.StatusCode())
	})

}

func TestPostComment(t *testing.T) {

	t.Run("cannot post comment with an Invalid Token", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			SetBody(`{"slug": "/", "author": "12345", "body": "hello world"}`).
			Post(BASE_URL + "/api/v1/comment")
		assert.NoError(t, err)

		assert.Equal(t, 401, resp.StatusCode())

	})

	t.Run("can post comment with Valid Token", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", "bearer "+CreateToken()).
			SetBody(`{"slug": "/", "author": "12345", "body": "hello world"}`).
			Post(BASE_URL + "/api/v1/comment")
		assert.NoError(t, err)

		assert.Equal(t, 200, resp.StatusCode())

	})

}
