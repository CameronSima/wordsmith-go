package game

// Service represents services for game
type Service interface {
	ChecksOut(g Game) bool
}

type service struct {
	repo Repository
}

// NewGameService returns a new Game Service
func NewGameService(repo Repository) Service {
	return &service{repo}
}

func (s *service) ChecksOut(g Game) bool {
	return true
}
