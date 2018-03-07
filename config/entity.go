package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"wordsmith-go/bonus"
)

// Levels contains all the levels available to a User
type Levels []LevelConfig

// LevelConfig represents a user's current level
type LevelConfig struct {
	Level             int           `json:"level"`
	PointsRequired    int           `json:"pointsRequired"`
	Bonuses           []bonus.Bonus `json:"bonuses"`
	NumVowels         int           `json:"numVowels"`
	NumLetters        int           `json:"numLetters"`
	PointsToNextLevel int           `json:"pointsToNextLevel"`
}

// NewLevelConfig returns a new instance of Levels[]
func NewLevelConfig() Levels {
	absPath, _ := filepath.Abs("../config/config.json")
	file, err := os.Open(absPath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	levels := Levels{}
	err = decoder.Decode(&levels)
	if err != nil {
		log.Fatal(err)
	}
	return levels
}