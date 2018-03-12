package game

import (
	"math/rand"
	"strings"
	"time"
)

// Word is what this game is all about!
type Word struct {
	Letters     []Letter `json:"letters"`
	Definitions []string `json:"definitions"`
	Score       int      `json:"score"`
}

func (l Letter) isVowel() bool {
	switch l.Value {
	case "A":
		return true
	case "E":
		return true
	case "I":
		return true
	case "O":
		return true
	case "U":
		return true
	default:
		return false
	}
}

// Return a single random definition
func (w Word) getDefinition() string {
	rand.Seed(time.Now().Unix())
	return w.Definitions[rand.Intn(len(w.Definitions))]
}

func (w *Word) getScore() int {
	if w.Score == 0 {
		sum := 0
		for _, letter := range w.Letters {
			sum += letter.Points
		}
		length := len(w.Letters)
		w.Score = Formulas[length][0]*sum + Formulas[length][1]
	}
	return w.Score
}

func (w Word) toString() string {
	wordStr := ""
	for _, letter := range w.Letters {
		wordStr += letter.Value
	}
	return wordStr
}

func (w Word) equals(w2 Word) bool {
	return w.toString() == w2.toString()
}

func (w Word) fromString(str string) {
	strArr := strings.Split(str, "")
	letters := make([]Letter, len(strArr))

	for i, letterStr := range strArr {
		value := strings.ToUpper(letterStr)
		letter := Letter{
			Value:  value,
			Points: Points[value],
		}
		letters[i] = letter
	}
	w.Letters = letters
}
