package game

import (
	"time"

	"wordsmith-go/bonus"
	"wordsmith-go/config"

	"google.golang.org/appengine/datastore"
)

// Game is what this is all about!
type Game struct {
	Ended       bool           `json:"ended"`
	Letterset   []Letter       `json:"letterset"`
	FinalScore  int            `json:"finalScore"`
	GameScore   int            `json:"gameScore" datastore:"-"`
	TimeBonus   int            `json:"timeBonus" datastore:"-"`
	TopWord     Word           `json:"topWord" datastore:"_"`
	WordsUsed   []Word         `json:"wordsUsed" datastore:"-"`
	StartTime   time.Time      `json:"startTime"`
	EndTime     time.Time      `json:"endTime"`
	BonusesUsed []bonus.Bonus  `json:"bonuses" datastore:"-"`
	UserKey     *datastore.Key `json:"userKey"`
	Key         *datastore.Key `json:"key"`
}

// NewGame returns a new game based on a user's level.
func NewGame(l config.LevelConfig) Game {
	numLetters := l.NumLetters
	numVowels := l.NumVowels
	ls := NewLetterSet(numLetters, numVowels)

	g := Game{
		Ended:     false,
		Letterset: ls,
		StartTime: time.Now(),
	}
	return g
}
