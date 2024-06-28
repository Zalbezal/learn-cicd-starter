package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/users", http.NoBody)
	recorder := httptest.NewRecorder()
	recorder.HeaderMap.Add("Authorization", "ApiKey =123123123123")

	authHeader := recorder.HeaderMap.Get("Authorization")
	if authHeader == "" {
		t.Fatalf("ErrNoAuthHeaderIncluded")
	}
	splitAuth := strings.Split(authHeader, " ")
	fmt.Println(splitAuth)
	if len(splitAuth) > 2 || splitAuth[0] != "ApiKey" {
		t.Fatal("wrong api key")
	}
	fmt.Println(request)
}

// // GetAPIKey -
// func GetAPIKey(headers http.Header) (string, error) {
// 	authHeader := headers.Get("Authorization")
// 	if authHeader == "" {
// 		return "", ErrNoAuthHeaderIncluded
// 	}
// 	splitAuth := strings.Split(authHeader, " ")
// 	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
// 		return "", errors.New("malformed authorization header")
// 	}

// 	return splitAuth[1], nil
// }
