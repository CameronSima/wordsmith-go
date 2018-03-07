package handlers

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"

	"wordsmith-go/game"
	"wordsmith-go/user"
)

type NewGameHandler struct {
	Repo user.Repository
}

func (h NewGameHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)

	// marshal to struct
	userReq := user.User{}
	err := json.NewDecoder(req.Body).Decode(&userReq)
	if err != nil {
		println("ERROR:")
		println(err.Error())
		http.Error(rw, "a valid user was not supplied", http.StatusBadRequest)
		return
	}

	println("REQES")
	println(userReq.Username)
	userRec, err := h.Repo.Find(c, &userReq)
	if err != nil {
		http.Error(rw, err.Error()+" could not find user", http.StatusInternalServerError)
		return
	}
	g := game.NewGame(userRec.LevelConfig)
	json.NewEncoder(rw).Encode(g)
}
