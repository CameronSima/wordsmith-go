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
	Score              int
}

// NewLetterSet returns a new letterset
func NewLetterSet(numLetters int, numVowels int) []Letter {
	lb := Letterset{
		numLetters: numLetters,
		numVowels:  numVowels,
	}
	return lb.build()
}

func (l *Letterset) build() []Letter {
	l.totalLetterWeights = l.getTotalWeights()

	vowelCount := 0
	letters := make([]Letter, l.numLetters)

	for vowelCount != l.numVowels {
		letters = l.getLetters(l.numLetters)
		vowelCount = l.getVowelCount(letters)
	}
	return letters
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

func (l *Letterset) scoreLetters() int {
	weights := l.totalLetterWeights
	f := float64(weights)
	scoreF := (1 - (0.001 * f)) * 70000
	return int(scoreF)

}
