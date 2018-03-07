package bonus

// Bonus determines the bonuses available to a user
type Bonus struct {
	Type  string `json:"type"`
	Level int    `json:"level"`
	Count int    `json:"count"`
}
