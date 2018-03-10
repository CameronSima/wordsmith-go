package game

import (
	"math/rand"
	"time"
)

// Letterset builds a letterset
type Letterset struct {
	Letters            []Letter
	totalLetterWeights int
	numLetters         int
	numVowels          int
	score              int
}

// implement sort()
func (l Letterset) Len() int { return len(l.Letters) }
func (l Letterset) Swap(i, j int) {
	l.Letters[i], l.Letters[j] = l.Letters[j], l.Letters[i]
}
func (ls Letterset) Less(i, j int) bool {
	return ls.Letters[i].Value < ls.Letters[j].Value
}

// NewLetterSet returns a new letterset
func NewLetterSet(numLetters int, numVowels int) Letterset {
	l := Letterset{
		numLetters: numLetters,
		numVowels:  numVowels,
	}
	l.build()
	return l
}

func (l *Letterset) build() {
	l.totalLetterWeights = l.getTotalWeights()

	vowelCount := 0
	letters := make([]Letter, l.numLetters)

	for vowelCount != l.numVowels {
		letters = l.getLetters(l.numLetters)
		vowelCount = l.getVowelCount(letters)
	}
	l.Letters = letters
}

func (l Letterset) getLetters(numLetters int) []Letter {
	letters := make([]Letter, numLetters)

	index := 0
	for index < numLetters {
		letter := l.getRandomLetter()
		letter.ID = index
		letters[index] = letter
		index++
	}
	return letters
}

// GetRandomLetter
func (l Letterset) getRandomLetter() Letter {
	var letter Letter
	upto := 0
	ri := l.getRandomInt()

	for key, val := range Weights {
		if upto+val > ri {
			letter = Letter{
				Value:       key,
				Points:      Points[key],
				Selected:    false,
				Highlighted: false,
			}
			break
		}
		upto += val
	}
	return letter
}

// GetVowelCount
func (l Letterset) getVowelCount(letters []Letter) int {
	count := 0
	for _, letter := range letters {
		if letter.isVowel() {
			count++
		}
	}
	return count
}

// GetRandomFloat between 0 and total weights
func (l Letterset) getRandomInt() int {
	max := l.totalLetterWeights
	rand.Seed(time.Now().UTC().Unix())
	return rand.Intn(max)
}

// GetTotalWeights of letters
func (l Letterset) getTotalWeights() int {
	total := 0
	for _, weight := range Weights {
		total += weight
	}
	return total
}

func (l *Letterset) Score() int {

	weights := 0
	for value, weight := range Weights {
		for _, lett := range l.Letters {
			if lett.Value == value {
				weights += weight
			}
		}
	}

	println(weights)
	f := float64(weights)
	scoreF := (1 - (0.000001 * f)) * 70000
	l.score = int(scoreF)

	println(l.score)
	return l.score
}
