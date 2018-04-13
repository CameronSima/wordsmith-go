package user

import (
	"testing"
	"wordsmith-go/bonus"

	"wordsmith-go/config"
	"wordsmith-go/game"
)

func TestGetUserLevel(t *testing.T) {
	confs := config.NewLevelConfig()

	u := User{
		GameStats: GameStats{
			Points: 500000,
		},
		LevelConfig: confs[0],
	}
	expected := 1
	actual := u.getUserLevel(confs).Level

	if expected != actual {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expected, actual)
	}
}

func TestGetUserLevel4(t *testing.T) {
	confs := config.NewLevelConfig()

	u := User{
		GameStats: GameStats{
			Points: 2013789,
		},
		LevelConfig: confs[0],
	}
	expected := 4
	actual := u.getUserLevel(confs).Level

	expectedPointsToNextLevel := 2500000
	actualPointsToNextLevel := u.getUserLevel(confs).PointsToNextLevel

	if expected != actual {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expected, actual)
	}

	if expectedPointsToNextLevel != actualPointsToNextLevel {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expectedPointsToNextLevel, actualPointsToNextLevel)
	}
}

func TestUpdateUserLevel(t *testing.T) {
	confs := config.NewLevelConfig()

	u := User{
		GameStats: GameStats{
			Points: 15000,
			TopWord: game.Word{
				Score: 10000,
			},
		},
		LevelConfig: confs[0],
	}

	g := game.Game{
		FinalScore: 1005000,
		TopWord: game.Word{
			Score: 20000,
		},
	}
	u.UpdateStats(confs, g)

	expectedLevel := 2
	actualLevel := u.getUserLevel(confs).Level
	if expectedLevel != actualLevel {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expectedLevel, actualLevel)
	}

	expectedTopWordScore := 20000
	actualTopWordScore := u.GameStats.TopWord.Score
	if expectedTopWordScore != actualTopWordScore {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expectedTopWordScore, actualTopWordScore)
	}

	expectedPoints := 1020000
	actualPoints := u.GameStats.Points
	if expectedPoints != actualPoints {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expectedTopWordScore, actualTopWordScore)
	}
}

func TestUpdateUserLevel3(t *testing.T) {
	confs := config.NewLevelConfig()

	u := User{
		GameStats: GameStats{
			Points: 1400000,
			TopWord: game.Word{
				Score: 10000,
			},
		},
		LevelConfig: confs[1],
	}

	g := game.Game{
		FinalScore: 200000,
		TopWord: game.Word{
			Score: 20000,
		},
	}
	u.UpdateStats(confs, g)

	expectedLevel := 3
	actualLevel := u.getUserLevel(confs).Level
	if expectedLevel != actualLevel {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expectedLevel, actualLevel)
	}
}

func TestMergeBonus(t *testing.T) {
	u := User{}
	letters := []bonus.Bonus{
		bonus.Bonus{
			Value: "A",
			Count: 2,
		},
		bonus.Bonus{
			Value: "B",
			Count: 1,
		},
	}
	bonuses := []bonus.Bonus{
		bonus.Bonus{
			Type:  "WordHintBonus",
			Count: 1,
		},
	}
	u.Letters = letters
	u.Bonuses = bonuses
	u.BonusSelectionPoints = 6

	purchasedBonuses := []bonus.Bonus{
		bonus.Bonus{
			Type:  "WordHintBonus",
			Count: 1,
		},
		bonus.Bonus{
			Type:  "LetterBonus",
			Count: 2,
		},
	}

	success := u.UpdateBonuses(purchasedBonuses)

	expectedWordHintBonuses := 2
	expectedLetterBonuses := 2
	expectedNumOfLetters := 5
	expectedRemainingBonusSelectionPoints := 0
	actualWordHintBonuses := 0
	actualLetterBonuses := 0
	actualNumOfLetters := 0
	actualRemainingBonusSelectionPoints := u.BonusSelectionPoints

	for _, l := range u.Bonuses {
		switch val := l.Type; val {
		case "WordHintBonus":
			actualWordHintBonuses += l.Count
		case "LetterBonus":
			actualLetterBonuses += l.Count
		}
	}

	for _, lb := range u.Letters {
		println(lb.Value)
		actualNumOfLetters += lb.Count
	}

	if success != true {
		t.Errorf("Test failed, merge was not successful (not enough selectionPoints)")
	}

	if expectedNumOfLetters != actualNumOfLetters {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expectedNumOfLetters, actualNumOfLetters)
	}

	if len(u.Letters) > 4 || len(u.Letters) < 2 {
		t.Errorf("Test failed, wrong number of letters")
	}

	if expectedWordHintBonuses != actualWordHintBonuses {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expectedWordHintBonuses, actualWordHintBonuses)
	}
	if expectedLetterBonuses != actualLetterBonuses {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expectedLetterBonuses, actualLetterBonuses)
	}
	if actualRemainingBonusSelectionPoints != actualRemainingBonusSelectionPoints {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expectedRemainingBonusSelectionPoints, actualRemainingBonusSelectionPoints)
	}

}

func TestMergeBonusFail(t *testing.T) {
	u := User{}
	letters := []bonus.Bonus{
		bonus.Bonus{
			Value: "A",
			Count: 2,
		},
		bonus.Bonus{
			Value: "B",
			Count: 1,
		},
	}
	bonuses := []bonus.Bonus{
		bonus.Bonus{
			Type:  "WordHintBonus",
			Count: 1,
		},
	}
	u.Letters = letters
	u.Bonuses = bonuses
	u.BonusSelectionPoints = 4

	purchasedBonuses := []bonus.Bonus{
		bonus.Bonus{
			Type:  "WordHintBonus",
			Count: 1,
		},
		bonus.Bonus{
			Type:  "LetterBonus",
			Count: 4,
		},
	}

	success := u.UpdateBonuses(purchasedBonuses)

	expectedLetterACount := 2
	expectedLetterBCount := 1
	expectedWordHintBonuses := 1
	expectedLetterBonuses := 0
	expectedRemainingBonusSelectionPoints := 4
	actualLetterACount := 0
	actualLetterBCount := 0
	actualWordHintBonuses := 0
	actualLetterBonuses := 0
	actualRemainingBonusSelectionPoints := u.BonusSelectionPoints

	for _, l := range u.Letters {
		switch val := l.Value; val {
		case "A":
			actualLetterACount += l.Count
		case "B":
			actualLetterBCount += l.Count
		default:
			println(val)
		}
	}

	for _, l := range u.Bonuses {
		switch val := l.Type; val {
		case "WordHintBonus":
			actualWordHintBonuses += l.Count
		case "LetterBonus":
			actualLetterBonuses += l.Count
		}
	}

	if success != false {
		t.Errorf("Test failed, merge was not successful (not enough selectionPoints)")
	}

	if expectedLetterACount != actualLetterACount {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expectedLetterACount, actualLetterACount)
	}
	if expectedLetterBCount != actualLetterBCount {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expectedLetterBCount, actualLetterBCount)
	}
	if expectedWordHintBonuses != actualWordHintBonuses {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expectedWordHintBonuses, actualWordHintBonuses)
	}
	if expectedLetterBonuses != actualLetterBonuses {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expectedLetterBonuses, actualLetterBonuses)
	}
	if actualRemainingBonusSelectionPoints != actualRemainingBonusSelectionPoints {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expectedRemainingBonusSelectionPoints, actualRemainingBonusSelectionPoints)
	}

}
