package game

import (
	"testing"
)

func TestGetTotalWeights(t *testing.T) {
	expected := 852894
	ls := NewLetterSet(8, 2)
	actual := ls.getTotalWeights()

	if actual != expected {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expected, actual)
	}
}

func TestGetRandomLetter(t *testing.T) {
	lb := Letterset{}
	lb.build()
	randomLetter := lb.getRandomLetter()

	t.Log(randomLetter.Value)
	for key, val := range Points {
		if randomLetter.Value == key {
			if randomLetter.Points != val {

				t.Errorf("Test failed, expected:'%d', got:'%d'", val, randomLetter.Points)
			}
		}
	}
}

func TestGetVowelCount(t *testing.T) {
	lb := Letterset{}
	lett1 := Letter{Value: "A"}
	lett2 := Letter{Value: "L"}
	lett3 := Letter{Value: "E"}
	lettArr := []Letter{lett1, lett2, lett3}

	expected := 2
	actual := lb.getVowelCount(lettArr)

	if expected != actual {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expected, actual)
	}
}

func TestScoreLetterSet(t *testing.T) {
	ls := NewLetterSet(6, 2)
	score := ls.Score()

	println(score)
}

func TestNewLetterset(t *testing.T) {
	ls := NewLetterSet(6, 2)
	expected := 6
	actual := len(ls.Letters)

	if expected != actual {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expected, actual)
	}
}
