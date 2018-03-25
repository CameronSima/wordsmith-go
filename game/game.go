package game

import (
	"time"

	"wordsmith-go/bonus"
	"wordsmith-go/config"
)

// Game is what this is all about!
type Game struct {
	NumWords       int           `json:"numWords" datastore:"-"`
	NumVowels      int           `json:"numVowels" datastore:"-"`
	Ended          bool          `json:"ended"`
	Letterset      []Letter      `json:"letterset" datastore:"-"`
	LettersetBonus int           `json:"lettersetBonus" datastore:"-"`
	FinalScore     int           `json:"finalScore"`
	GameScore      int           `json:"gameScore" datastore:"-"`
	GameTime       int           `json:"gameTime" datastore:"-"`
	TimeBonus      int           `json:"timeBonus" datastore:"-"`
	TopWord        Word          `json:"topWord" datastore:"-"`
	WordsUsed      []Word        `json:"wordsUsed" datastore:"-"`
	StartTime      time.Time     `json:"startTime"`
	EndTime        time.Time     `json:"endTime"`
	BonusesUsed    []bonus.Bonus `json:"bonuses" datastore:"-"`
	Username       string        `json:"username"`
}

// NewGame returns a new game based on a user's level.
func NewGame(l config.LevelConfig, username string) Game {
	numLetters := l.NumLetters
	numVowels := l.NumVowels
	ls := NewLetterSet(numLetters, numVowels)

	return Game{
		Ended:          false,
		Letterset:      ls.Letters,
		LettersetBonus: ls.Score(),
		StartTime:      time.Now(),
		Username:       username,
		GameTime:       l.GameTime,
		NumWords:       l.NumWords,
	}
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
