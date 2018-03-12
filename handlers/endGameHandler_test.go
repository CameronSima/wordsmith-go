package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"wordsmith-go/config"
	"wordsmith-go/game"
	"wordsmith-go/user"
)

func TestEndGame(t *testing.T) {

	gameRepo := game.NewTestRepository()
	userRepo := user.NewTestRepository()
	configs := config.NewLevelConfig()
	handler := EndGameHandler{
		GameRepo:     gameRepo,
		UserRepo:     userRepo,
		LevelConfigs: configs,
	}

	g := game.Game{
		Username: "cameron",
	}
	u := user.User{
		Username: "cameron",
		Password: "pass",
	}
	reqBody := EndGameRequest{
		User: u,
		Game: g,
	}

	b, _ := json.Marshal(reqBody)
	reader := bytes.NewReader(b)

	req := httptest.NewRequest("POST", "http://localhost:8080/", reader)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println("RESPONSE")
	fmt.Println(resp.StatusCode)
	println(string(body))

	expectedResponse := 200
	actualResponse := resp.StatusCode

	if expectedResponse != actualResponse {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expectedResponse, actualResponse)
	}
	return
}

func TestEndGameSubmittedEndedGame(t *testing.T) {

	gameRepo := game.NewTestRepository()
	userRepo := user.NewTestRepository()
	configs := config.NewLevelConfig()
	handler := EndGameHandler{
		GameRepo:     gameRepo,
		UserRepo:     userRepo,
		LevelConfigs: configs,
	}

	g := game.Game{
		Username: "cameron",
		Ended:    true,
	}
	u := user.User{
		Username: "cameron",
		Password: "pass",
	}
	reqBody := EndGameRequest{
		User: u,
		Game: g,
	}

	b, _ := json.Marshal(reqBody)
	reader := bytes.NewReader(b)

	req := httptest.NewRequest("POST", "http://localhost:8080/", reader)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	resp := w.Result()
	//body, _ := ioutil.ReadAll(resp.Body)

	expectedResponse := 500
	actualResponse := resp.StatusCode

	if expectedResponse != actualResponse {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expectedResponse, actualResponse)
	}
}
