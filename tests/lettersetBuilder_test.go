package tests

import (
	"testing"

	"appengine-guestbook-go/game"
)

func TestGetTotalWeights(t *testing.T) {
	expected := 852894
	lb := game.Letterset{}
	actual := lb.getTotalWeights()

	if actual != expected {
		t.Errorf("Test failed, expected:'%f', got:'%f'", expected, actual)
	}
}

// func TestGetRandomInt(t *testing.T) {
// 	lb := game.LettersetBuilder{}
// 	rf := lb.GetRandomInt()

// 	t.Log(rf)

// 	if rf > 0.999999 || rf < 0.0 {
// 		t.Errorf("Test failed,  got:'%f'", rf)
// 	}
// }

func TestGetRandomLetter(t *testing.T) {
	lb := game.LettersetBuilder{}
	randomLetter := lb.GetRandomLetter()

	t.Log(randomLetter.Value)
	for key, val := range game.Points {
		if randomLetter.Value == key {
			if randomLetter.Points != val {

				t.Errorf("Test failed, expected:'%d', got:'%d'", val, randomLetter.Points)
			}
		}
	}
}

func TestGetVowelCount(t *testing.T) {
	lb := game.LettersetBuilder{}
	lett1 := game.Letter{Value: "A"}
	lett2 := game.Letter{Value: "L"}
	lett3 := game.Letter{Value: "E"}
	lettArr := []game.Letter{lett1, lett2, lett3}

	expected := 2
	actual := lb.GetVowelCount(lettArr)

	if expected != actual {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expected, actual)
	}
}

func TestNewLetterset(t *testing.T) {
	ls := game.NewLetterSet(6, 2)
	expected := 6
	actual := len(ls)

	if expected != actual {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expected, actual)
	}
}
