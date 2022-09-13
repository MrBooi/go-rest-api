//go:build end_two_end
// +build end_two_end

package test

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

const (
	BASE_URL = "http://localhost:8080"
)

func CreateToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokeString, err := token.SignedString([]byte("FlutterGoDeveloper"))

	if err != nil {
		fmt.Println(err)
	}

	return tokeString
}
