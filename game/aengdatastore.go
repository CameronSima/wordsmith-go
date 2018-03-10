package game

import (
	"context"

	"google.golang.org/appengine/datastore"
)

type repo struct {
}

// NewDatastoreRepository return a new test repository
func NewDatastoreRepository() Repository {
	return &repo{}
}

func (r *repo) FindAll(ctx context.Context) (*[]Game, error) {
	q := datastore.NewQuery("User")

	games := make([]Game, 0)
	if _, err := q.GetAll(ctx, &games); err != nil {
		return &games, err
	}
	return &games, nil
}

func (r *repo) Find(ctx context.Context, g *Game) (*Game, error) {
	gameRec := Game{}

	keyString := g.StartTime.String() + g.Username
	key := datastore.NewKey(ctx, "Game", keyString, 0, nil)
	if err := datastore.Get(ctx, key, &gameRec); err != nil {
		return &gameRec, err
	}
	return &gameRec, nil
}

func (r *repo) Save(ctx context.Context, g *Game) (*Game, error) {
	keyString := g.StartTime.String() + g.Username
	key := datastore.NewKey(ctx, "Game", keyString, 0, nil)

	_, err := datastore.Put(ctx, key, g)
	if err != nil {
		return g, err
	}
	return g, nil
}
