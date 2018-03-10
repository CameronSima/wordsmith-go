package game

import "context"

// Repository is the interface for User DAO
type Repository interface {
	FindAll(ctx context.Context) (*[]Game, error)
	Find(ctx context.Context, user *Game) (*Game, error)
	Save(ctx context.Context, user *Game) (*Game, error)
}
