package config

import "testing"

func TestInit(t *testing.T) {
	expected := 1
	c := NewLevelConfig()
	actual := c[0].Level

	if actual != expected {
		t.Errorf("Test failed, expected:'%f', got:'%f'", expected, actual)
	}

	expected2 := "WordHintBonus"
	actual2 := c[1].Bonuses[0].Type

	if actual2 != expected2 {
		t.Errorf("Test failed, expected:'%s', got:'%s'", expected2, actual2)
	}
}
