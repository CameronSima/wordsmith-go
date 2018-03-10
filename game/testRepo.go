package game

import (
	"context"
)

type test_repo struct {
	savedGames []Game
}

// NewTestRepository return a new Datastore repository
func NewTestRepository() Repository {
	return &test_repo{}
}

func (r *test_repo) FindAll(ctx context.Context) (*[]Game, error) {
	g1 := Game{}
	g2 := Game{}
	g3 := Game{}
	games := []Game{g1, g2, g3}
	return &games, nil
}

func (r *test_repo) Find(ctx context.Context, g *Game) (*Game, error) {
	return g, nil
}

func (r *test_repo) Save(ctx context.Context, g *Game) (*Game, error) {
	return g, nil
}
