package user

import (
	"testing"

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
		t.Errorf("Test failed, expected:'%s', got:'%s'", expected, actual)
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

func TestAddBonus(t *testing.T) {
	confs := config.NewLevelConfig()

	u := User{
		GameStats: GameStats{
			Points: 15000,
			TopWord: game.Word{
				Score: 10000,
			},
		},
		LevelConfig: confs[1],
	}
	u.AddBonus("TimeBonus", 5)

	expected := 6
	actual := u.LevelConfig.Bonuses[1].Count

	if expected != actual {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expected, actual)
	}
}

func TestMergeBonus(t *testing.T) {
	confs := config.NewLevelConfig()

	u := User{
		GameStats: GameStats{
			Points: 15000,
			TopWord: game.Word{
				Score: 10000,
			},
		},
		LevelConfig: confs[1],
	}

	// user had an awesome game, scoring 1,005,000 points. Now should level up to
	// level 2, and has earned an additional 2 WordHint bonuses.
	g := game.Game{
		FinalScore: 1005000,
		TopWord: game.Word{
			Score: 20000,
		},
	}

	// simulate user purchase of 2 WordHint bonuses :)
	u.AddBonus("WordHintBonus", 2)

	u.UpdateStats(confs, g)

	expectedWordHintBonusCount := 3
	actualWordHintBonusCount := u.LevelConfig.Bonuses[0].Count

	if expectedWordHintBonusCount != actualWordHintBonusCount {
		t.Errorf("Test failed, expected:'%d', got:'%d'", expectedWordHintBonusCount, actualWordHintBonusCount)
	}

}
