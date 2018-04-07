package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"wordsmith-go/bonus"
	"wordsmith-go/user"
)

func TestUpdateBonuses(t *testing.T) {

	userRepo := user.NewTestRepository()
	handler := UpdateBonusesHandler{
		Repo: userRepo,
	}
	u := user.User{
		Username:             "cameron",
		Password:             "pass",
		BonusSelectionPoints: 10,
	}
	u.Bonuses = []bonus.Bonus{
		bonus.Bonus{
			Type:  "WordHintBonus",
			Count: 1,
		},
		bonus.Bonus{
			Type:  "TimeBonus",
			Count: 2,
		},
	}
	u.Letters = []bonus.Bonus{
		bonus.Bonus{
			Value: "A",
			Count: 1,
		},
		bonus.Bonus{
			Value: "E",
			Count: 2,
		},
	}

	newBonuses := []bonus.Bonus{
		bonus.Bonus{
			Type:  "WordHintBonus",
			Count: 1,
		},
		bonus.Bonus{
			Type:  "HighlightWordBonus",
			Count: 1,
		},
	}
	newLetters := []bonus.Bonus{
		bonus.Bonus{
			Value: "A",
			Count: 3,
		},
		bonus.Bonus{
			Value: "R",
			Count: 2,
		},
	}

	reqBody := updateBonusesRequest{
		user:    u,
		bonuses: newBonuses,
		letters: newLetters,
	}

	b, _ := json.Marshal(reqBody)
	reader := bytes.NewReader(b)

	req := httptest.NewRequest("POST", "/update-bonuses", reader)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	println(u.Username)
	println(string(body))

	expectedResponse := 200
	actualResponse := resp.StatusCode

	if expectedResponse != actualResponse {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expectedResponse, actualResponse)
	}

	return
}
