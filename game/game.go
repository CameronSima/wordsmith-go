package game

import (
	"time"

	"wordsmith-go/bonus"
	"wordsmith-go/config"

	"google.golang.org/appengine/datastore"
)

// Game is what this is all about!
type Game struct {
	Ended          bool           `json:"ended"`
	Letterset      []Letter       `json:"letterset"`
	LettersetBonus int            `json:"lettersetBonus"`
	FinalScore     int            `json:"finalScore"`
	GameScore      int            `json:"gameScore" datastore:"-"`
	TimeBonus      int            `json:"timeBonus" datastore:"-"`
	TopWord        Word           `json:"topWord" datastore:"_"`
	WordsUsed      []Word         `json:"wordsUsed" datastore:"-"`
	StartTime      time.Time      `json:"startTime"`
	EndTime        time.Time      `json:"endTime"`
	BonusesUsed    []bonus.Bonus  `json:"bonuses" datastore:"-"`
	Username       string         `json:"username"`
	Key            *datastore.Key `json:"key"`
}

// NewGame returns a new game based on a user's level.
func NewGame(l config.LevelConfig, username string) Game {
	numLetters := l.NumLetters
	numVowels := l.NumVowels
	ls := NewLetterSet(numLetters, numVowels)

	g := Game{
		Ended:          false,
		Letterset:      ls.Letters,
		LettersetBonus: ls.Score(),
		StartTime:      time.Now(),
		Username:       username,
	}
	return g
}

// ChecksOut checks a game for vailidity
func (submitted Game) ChecksOut(saved Game) bool {
	if submitted.Key != saved.Key {
		return false
	}

	//TODO: Add more checks! Check bonus, letters, etc.
	return true
}

//CheckScore checks a game's submitted score
func (g Game) CheckScore() bool {
	score := 0
	wordsScore := 0

	for _, word := range g.WordsUsed {
		wordsScore += word.getScore()
	}
	score += wordsScore
	score += g.TimeBonus
	score += g.GameScore
	score += g.LettersetBonus
	return score == g.FinalScore
}
