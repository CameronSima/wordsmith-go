package handlers

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"

	"wordsmith-go/user"
)

type UpdatePasswordRequest struct {
	User        user.User `json:"user"`
	NewPassword string    `json:"newPassword"`
}

type ChangePasswordHandler struct {
	Repo user.Repository
}

func (h ChangePasswordHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)

	// decode req
	updateReq := UpdatePasswordRequest{}
	err := json.NewDecoder(req.Body).Decode(&updateReq)
	if err != nil {
		http.Error(rw, "Could not decode the user", http.StatusInternalServerError)
		return
	}

	userRec, err := h.Repo.Find(c, &updateReq.User)
	if err != nil {
		http.Error(rw, err.Error()+" Could not find user", http.StatusInternalServerError)
		return
	}

	// check password
	err = userRec.CheckPassword(userRec.Password)
	if err != nil {
		http.Error(rw, "Invalid password", http.StatusBadRequest)
		return
	}

	hashed, err := user.HashPassword(updateReq.NewPassword)
	if err != nil {
		http.Error(rw, err.Error()+" error hashing password", http.StatusInternalServerError)
		return
	}

	userRec.Password = hashed

	updatedUser, err := h.Repo.Save(c, userRec)
	if err != nil {
		http.Error(rw, err.Error()+" couldn't save the updated user", http.StatusInternalServerError)
		return
	}

	// send response
	responseJSON, err := NewUserResponseJSON(updatedUser, "success")
	if err != nil {
		http.Error(rw, err.Error()+" error creating response object", http.StatusBadRequest)
		return
	}

	defer req.Body.Close()
	rw.Write(responseJSON)
	return
}
