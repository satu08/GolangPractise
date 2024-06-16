package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiError struct {
	Err    string
	Status int
}

func (e ApiError) Error() string {
	return e.Err
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if e, ok := err.(ApiError); ok {
				writeJson(w, e.Status, e)
				return
			}
			writeJson(w, http.StatusInternalServerError, ApiError{Err: "internal server", Status: http.StatusInternalServerError})
		}
	}
}

func main() {
	fmt.Println("satyanarayan Jadhav")
	http.HandleFunc("/user", makeHTTPHandleFunc(handleGetUserById))
	http.ListenAndServe(":3000", nil)
}

type user struct {
	ID    int
	Valid bool
}

func handleGetUserById(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return ApiError{Err: "invalid method", Status: http.StatusMethodNotAllowed}
	}
	user := user{}
	if !user.Valid {
		return ApiError{Err: "user not allowed", Status: http.StatusForbidden}
	}

	return writeJson(w, http.StatusOK, user)
}

func writeJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
