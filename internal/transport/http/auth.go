package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func JWTAuth(
	original func(w http.ResponseWriter, r *http.Request),
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]

		if authHeader == nil {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		autHeadersParts := strings.Split(authHeader[0], " ")
		if len(autHeadersParts) != 2 || strings.ToLower(autHeadersParts[0]) != "bearer" {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		if validateToken(autHeadersParts[1]) {
			original(w, r)
		} else {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

	}
}

func validateToken(accessToken string) bool {
	var mySigningkey = []byte("FlutterGoDeveloper")
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not validate auth token")
		}
		return mySigningkey, nil
	})

	if err != nil {
		return false
	}

	return token.Valid

}
