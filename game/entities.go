package game

// Formulas used to calculate word sore by length
var Formulas = [][]int{{1, 0}, {20, 2000},
	{70, 7000}, {80, 8000},
	{100, 10000}, {120, 12000},
	{140, 15000}, {180, 20000},
	{220, 25000}, {260, 30000},
	{350, 40000}, {440, 50000}}

// Points assigned to each letter
var Points = map[string]int{
	"A": 1, "B": 3, "C": 3, "D": 2, "E": 1,
	"F": 4, "G": 2, "H": 4, "I": 1, "J": 8,
	"K": 5, "L": 1, "M": 3, "N": 1, "O": 1, "P": 3,
	"Q": 10, "R": 1, "S": 1, "T": 1, "U": 1,
	"V": 4, "W": 4, "X": 8, "Y": 4, "Z": 10,
}

//Weights are a letter's probability of appearing
var Weights = map[string]int{
	"A": 79971, "B": 20199, "C": 43593,
	"D": 32184, "E": 112513, "F": 14622,
	"G": 22845, "H": 23497, "I": 8563,
	"J": 1693, "K": 8535, "L": 60585,
	"M": 28389, "N": 71596, "O": 65476,
	"P": 29165, "Q": 183, "R": 7264,
	"S": 66044, "T": 70926, "U": 37091,
	"V": 11369, "W": 8899, "X": 2925,
	"Y": 24432, "Z": 335,
}

// Dictionary is what the game uses to check words
type Dictionary struct {
	Entries map[string]map[string][]string
}

// Letter is the basic building block of a Word.
type Letter struct {
	Value       string `json:"value"`
	ID          string `json:"id"`
	Points      int    `json:"points"`
	Weight      int    `json:"-"`
	Selected    bool   `json:"-"`
	Highlighted bool   `json:"-"`
}
