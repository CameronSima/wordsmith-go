package handlers

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"

	"wordsmith-go/game"
	"wordsmith-go/user"
)

type NewGameHandler struct {
	UserRepo user.Repository
	GameRepo game.Repository
}

func (h NewGameHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)

	// marshal to struct
	userReq := user.User{}
	err := json.NewDecoder(req.Body).Decode(&userReq)
	if err != nil {
		http.Error(rw, "a valid user was not supplied", http.StatusBadRequest)
		return
	}

	userRec, err := h.UserRepo.Find(c, &userReq)
	if err != nil {
		http.Error(rw, err.Error()+" could not find user", http.StatusInternalServerError)
		return
	}
	g := game.NewGame(userRec.LevelConfig, userRec.Username)
	_, err = h.GameRepo.Save(c, &g)
	if err != nil {
		http.Error(rw, err.Error()+" could not save the game", http.StatusInternalServerError)
	}

	json.NewEncoder(rw).Encode(g)
}
