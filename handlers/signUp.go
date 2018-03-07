package handlers

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"

	"wordsmith-go/config"
	"wordsmith-go/user"
)

// SignUpHandler handles a signUp REST call
type SignUpHandler struct {
	Repo         user.Repository
	LevelConfigs config.Levels
}

func (h SignUpHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)

	// marshal to struct
	userReq := user.User{}
	err := json.NewDecoder(req.Body).Decode(&userReq)
	if err != nil {
		http.Error(rw, "invalid user data supplied", http.StatusInternalServerError)
		return
	}

	// check user doesn't already exist
	existingUser, err := h.Repo.Find(c, &userReq)
	if err == nil || existingUser.Username == userReq.Username {
		http.Error(rw, "user already exists with that username", http.StatusBadRequest)
		return
	}

	// Add default values and hash password
	u, err := user.NewUser(userReq, h.LevelConfigs)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save user
	userRec, err := h.Repo.Save(c, &u)
	if err != nil {
		http.Error(rw, err.Error()+"heres the err", http.StatusInternalServerError)
		return
	}

	// send response
	responseJSON, err := NewUserResponseJSON(userRec, "success")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	defer req.Body.Close()
	rw.Write(responseJSON)
	return
}
