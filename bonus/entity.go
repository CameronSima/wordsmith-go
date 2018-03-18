package bonus

// Bonus determines the bonuses available to a user
type Bonus struct {
	Type  string `json:"type"`
	Level int    `json:"level"`
	Count int    `json:"count"`
}

// LetterBonus describes the value and count of a user's bonus letters
type LetterBonus struct {
	Value string `json:"value"`
	Count int    `json:"count"`
}
