package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"wordsmith-go/user"
)

type UserResponse struct {
	Status string     `json:"status"`
	Data   *user.User `json:"data"`
}

//NewUserResponseJSON takes a user and a status and return a json byte array
func NewUserResponseJSON(user *user.User, status string) ([]byte, error) {

	// remove password from response
	user.Password = ""

	userResponse := UserResponse{
		Status: status,
		Data:   user,
	}

	responseJSON, err := json.Marshal(userResponse)
	if err != nil {
		return []byte{}, errors.New("couldn't marshal the user response")
	}
	return responseJSON, nil
}

// CorsHandler is a middleware for handling cors.
func CorsHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type,OPTIONS,GET,PUT,POST")
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("200 -- OK"))
		} else {
			h.ServeHTTP(w, r)
		}
		return
	}
}
