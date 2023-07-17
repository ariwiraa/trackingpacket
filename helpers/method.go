package helpers

import (
	"errors"
	"net/http"
)

func IsPost(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return errors.New("method not allowed")
	}

	return nil
}

func IsGet(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return errors.New("method not allowed")
	}

	return nil
}

func IsPut(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPut {
		return errors.New("method not allowed")
	}

	return nil
}
