package bonus

// Bonus determines the bonuses available to a user
type Bonus struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
	Count int    `json:"count"`
}
