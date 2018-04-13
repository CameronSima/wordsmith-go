package handlers

import (
	"encoding/json"
	"net/http"
	"wordsmith-go/bonus"
	"wordsmith-go/user"

	"google.golang.org/appengine"
)

// UpdateBonusesHandler is the endpoint for a user to purchase bonuses
type UpdateBonusesHandler struct {
	Repo user.Repository
}

// UpdateBonusesRequest represents the structure of a
// request to update a user's bonuses
type UpdateBonusesRequest struct {
	User    user.User     `json:"user"`
	Bonuses []bonus.Bonus `json:"bonuses"`
	Letters []bonus.Bonus `json:"letters"`
}

// ServeHTTP is what handles the request
func (h UpdateBonusesHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	// marshal to struct
	updateReq := UpdateBonusesRequest{}
	err := json.NewDecoder(req.Body).Decode(&updateReq)
	if err != nil {
		http.Error(rw, err.Error()+", couldn't decode the request", http.StatusInternalServerError)
		return
	}

	// Find user
	userRec, err := h.Repo.Find(ctx, &updateReq.User)
	if err != nil {
		http.Error(rw, err.Error()+" Could not find user", http.StatusInternalServerError)
		return
	}

	// check that the user has enough points to get the bonuses,
	// and subtract from total.
	updated := userRec.UpdateBonuses(updateReq.Bonuses)
	if updated == false {
		responseJSON, err := NewUserResponseJSON(userRec, "success")
		if err != nil {
			http.Error(rw, err.Error()+" couldn't create the user response JSON", http.StatusInternalServerError)
		}
		rw.Write(responseJSON)
		return
	}

	updatedUser, err := h.Repo.Save(ctx, userRec)
	if err != nil {
		http.Error(rw, err.Error()+"could not save the user", http.StatusInternalServerError)
		return
	}

	responseJSON, err := NewUserResponseJSON(updatedUser, "success")
	if err != nil {
		http.Error(rw, err.Error()+" couldn't create the user response JSON", http.StatusInternalServerError)
	}

	defer req.Body.Close()
	rw.Write(responseJSON)
	return
}
