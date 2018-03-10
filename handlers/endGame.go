package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"google.golang.org/appengine"

	"wordsmith-go/config"
	"wordsmith-go/game"
	"wordsmith-go/user"
)

// EndGameHandler handles a end of game request. Expects a user and game populated
// with appropriate results.
type EndGameHandler struct {
	GameRepo     game.Repository
	UserRepo     user.Repository
	LevelConfigs config.Levels
}

func (h EndGameHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)
	defer req.Body.Close()

	// decode req
	endGameRequest := EndGameRequest{}
	err := json.NewDecoder(req.Body).Decode(&endGameRequest)
	if err != nil {
		http.Error(rw, "Could not process the request", http.StatusInternalServerError)
		return
	}
	gameReq := endGameRequest.Game
	userReq := endGameRequest.User

	//VALIDATE USER
	// Find the user that submitted the game
	u, err := h.UserRepo.Find(c, &userReq)
	if err != nil {
		http.Error(rw, err.Error()+" Could not load user associated with this game", http.StatusInternalServerError)
		return
	}

	if err = u.CheckPassword(userReq.Password); err != nil {
		http.Error(rw, err.Error()+" Incorrect user password", http.StatusInternalServerError)
		return
	}

	// Find the submitted game from db
	gameRec, err := h.GameRepo.Find(c, &gameReq)
	if err != nil {
		http.Error(rw, err.Error()+" Could not find that game", http.StatusInternalServerError)
		return
	}

	//TODO: check game
	if gameRec.Ended == true {
		http.Error(rw, "the submitted game has already been completed", http.StatusInternalServerError)
		return
	}

	// If all is well, update the game and save in db
	gameReq.Ended = true
	gameReq.EndTime = time.Now()

	updatedGame, err := h.GameRepo.Save(c, &gameReq)
	if err != nil {
		http.Error(rw, err.Error()+" Error saving game", http.StatusInternalServerError)
		return
	}

	u.UpdateStats(h.LevelConfigs, *updatedGame)
	// send response
	responseJSON, err := NewUserResponseJSON(u, "end game success")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.Write(responseJSON)
	return
}
