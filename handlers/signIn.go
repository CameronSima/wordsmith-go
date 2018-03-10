package handlers

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"

	"wordsmith-go/user"
)

type SignInHandler struct {
	Repo user.Repository
}

func (h SignInHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)

	// decode req
	userReq := user.User{}
	err := json.NewDecoder(req.Body).Decode(&userReq)
	if err != nil {
		http.Error(rw, "Could not process the request", http.StatusInternalServerError)
		return
	}

	userRec, err := h.Repo.Find(c, &userReq)
	if err != nil {
		http.Error(rw, err.Error()+" Could not find user", http.StatusInternalServerError)
		return
	}

	// check password
	err = userRec.CheckPassword(userReq.Password)
	if err != nil {
		http.Error(rw, "Invalid password", http.StatusBadRequest)
		return
	}

	userRec.Password = ""

	// marshal back to json
	userJSON, err := json.Marshal(userRec)
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}

	defer req.Body.Close()
	rw.Write(userJSON)
	return
}
