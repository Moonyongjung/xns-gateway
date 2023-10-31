package util

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	unsufficientErr = "unsufficient data in the API"
)

func ParseAPI1Data(r *http.Request) (string, error) {
	api := r.URL.Path
	splited := strings.Split(api, "/")
	if len(splited) < 5 {
		return "", fmt.Errorf(unsufficientErr)
	}
	data := splited[4]
	return data, nil
}

func ParseAPI2Data(r *http.Request) (string, error) {
	api := r.URL.Path
	splited := strings.Split(api, "/")
	if len(splited) < 6 {
		return "", fmt.Errorf(unsufficientErr)
	}
	data := splited[5]
	return data, nil
}

func ParseAPI3Data(r *http.Request) (string, error) {
	api := r.URL.Path
	splited := strings.Split(api, "/")
	if len(splited) < 7 {
		return "", fmt.Errorf(unsufficientErr)
	}
	data := splited[6]
	return data, nil
}
