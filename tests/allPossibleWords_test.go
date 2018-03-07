package tests

import (
	"testing"

	"appengine-guestbook-go/game"
)

func TestGetAllPossibleWords(t *testing.T) {
	dict := game.Dictionary{}
	dict.Build()
	ls := []game.Letter{
		game.Letter{
			Value: "E",
		},
		game.Letter{
			Value: "A",
		},
		game.Letter{
			Value: "R",
		},
		game.Letter{
			Value: "L",
		},
		game.Letter{
			Value: "E",
		},
		game.Letter{
			Value: "T",
		},
		game.Letter{
			Value: "N",
		},
		game.Letter{
			Value: "L",
		},
	}

	apw := game.GetAllPossibleWords(ls, &dict)
	expected := 82
	actual := len(apw)

	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got:'%d'", expected, actual)
	}
}
