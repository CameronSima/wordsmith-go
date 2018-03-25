package guestbook

import (
	"context"
	"encoding/json"
	"html/template"
	"net/http"
	"wordsmith-go/game"

	"wordsmith-go/config"

	"wordsmith-go/handlers"
	"wordsmith-go/oauth"
	"wordsmith-go/user"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func init() {

	// Initialize level configs and repos
	configs := config.NewLevelConfig()
	userRepo := user.NewDatastoreRepository()
	gameRepo := game.NewDatastoreRepository()

	// initialize handlers
	signUpHandler := handlers.SignUpHandler{
		Repo:         userRepo,
		LevelConfigs: configs,
	}
	SignInHandler := handlers.SignInHandler{
		Repo: userRepo,
	}
	newGameHandler := handlers.NewGameHandler{
		UserRepo: userRepo,
		GameRepo: gameRepo,
	}
	endGameHandler := handlers.EndGameHandler{
		UserRepo:     userRepo,
		GameRepo:     gameRepo,
		LevelConfigs: configs,
	}

	http.HandleFunc("/", root)
	http.HandleFunc("/login", handlers.CorsHandler(SignInHandler))
	http.HandleFunc("/endGame", handlers.CorsHandler(endGameHandler))
	http.Handle("/signUp", handlers.CorsHandler(signUpHandler))
	http.HandleFunc("/allUsers", allUsers)
	http.Handle("/newGame", handlers.CorsHandler(newGameHandler))

	http.HandleFunc("/facebook_login", oauth.HandleFacebookLogin)
}

func allUsers(rw http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)
	q := datastore.NewQuery("User").Ancestor(userKey(c))

	users := make([]user.User, 0)
	if _, err := q.GetAll(c, &users); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(rw).Encode(users)
}

// Awesome landing page
func root(rw http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("public/index.html")
	t.Execute(rw, nil)
	return
}

func userKey(c context.Context) *datastore.Key {
	return datastore.NewKey(c, "User", "default_user", 0, nil)
}
