package user

import (
	"errors"
	"wordsmith-go/bonus"
	"wordsmith-go/config"
	"wordsmith-go/game"

	"golang.org/x/crypto/bcrypt"
)

// User represents a user.
type User struct {
	Username             string             `json:"username"`
	Password             string             `json:"password"`
	Email                string             `json:"email"`
	LevelConfig          config.LevelConfig `json:"levelConfig"`
	GameStats            GameStats          `json:"gameStats"`
	IsAdmin              bool               `json:"-"`
	BonusSelectionPoints int                `json:"bonusSelectionPoints"`

	// Letters are stored separately on the user entity because appengine doesn't allow
	// arrays of arrays ([[]LetterBonus]Bonus)
	Letters []bonus.Bonus `json:"letters"`
	Bonuses []bonus.Bonus `json:"bonuses"`
}

// GameStats are a user's Game stats, such as top word.
type GameStats struct {
	TopWord   game.Word `json:"topWord"`
	Points    int       `json:"points"`
	HighScore int       `json:"highScore"`
}

// NewUser returns a new User with defaults and hashed password.
func NewUser(u User, configs config.Levels) (User, error) {
	result := User{}
	hashed, err := HashPassword(u.Password)
	if err != nil {
		return result, err
	}

	// testing
	l := []bonus.Bonus{
		bonus.Bonus{
			Value: "A",
			Count: 1,
		},
		bonus.Bonus{
			Value: "E",
			Count: 2,
		},
		bonus.Bonus{
			Value: "R",
			Count: 3,
		},
	}
	result.Letters = l

	result.LevelConfig = configs[0]
	result.Username = u.Username
	result.Password = hashed
	result.IsAdmin = false
	return result, nil
}

// UpdateStats ensures a User's LevelConfig and GameStats reflect his current point count
// and levels him up if appropriate.
func (u *User) UpdateStats(l config.Levels, g game.Game) {
	u.updateGameStats(g)
	userLevel := u.getUserLevel(l)
	if userLevel.Level == u.LevelConfig.Level {
		return
	}
	u.LevelConfig = userLevel
}

func (u *User) updateGameStats(g game.Game) {
	u.GameStats.Points += g.FinalScore
	u.updatedHighScore(g)
	u.updateTopWord(g)
}

func (u *User) updateTopWord(g game.Game) {
	if g.TopWord.Score > u.GameStats.TopWord.Score {
		u.GameStats.TopWord = g.TopWord
	}
}

func (u *User) updatedHighScore(g game.Game) {
	if g.FinalScore > u.GameStats.HighScore {
		u.GameStats.HighScore = g.FinalScore
	}
}

func (u User) getUserLevel(l config.Levels) config.LevelConfig {
	var userLevel config.LevelConfig
	userPoints := u.GameStats.Points

	// iterate backwards
	for i := len(l) - 1; i >= 0; i-- {
		config := l[i]
		if userPoints > config.PointsRequired {
			userLevel = config
			break
		}
	}
	return userLevel
}

// CheckPassword checks a user's saved hashed password against a string.
func (u User) CheckPassword(pw string) error {
	hashedPwBtyeArr := []byte(u.Password)
	pwBtyeArr := []byte(pw)
	return bcrypt.CompareHashAndPassword(hashedPwBtyeArr, pwBtyeArr)
}

// HashPassword turns a string into a hash for password storage
func HashPassword(pw string) (string, error) {
	byteArr := []byte(pw)
	hashed, err := bcrypt.GenerateFromPassword(byteArr, bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("Couldn't hash password")
	}
	return string(hashed), nil
}

// UpdateBonuses Adds bonuses to a User struct. Returns true if the transaction
// was made, and false if the user doesn't have enough points.
func (u *User) UpdateBonuses(bonuses []bonus.Bonus) bool {

	// check user has enough points
	hasEnoughPoints := u.checkPoints(bonuses)
	if !hasEnoughPoints {
		return false
	}
	u.mergeBonuses(bonuses)
	u.buildLetters(bonuses)
	return true
}

func (u *User) checkPoints(bonuses []bonus.Bonus) bool {
	debitTotal := 0
	for _, b := range bonuses {
		debitTotal += b.Count
	}
	if debitTotal > u.BonusSelectionPoints {
		return false
	}
	u.BonusSelectionPoints -= debitTotal
	return true
}

func (u *User) mergeBonuses(bonuses []bonus.Bonus) {
	result := make(map[string]int)
	merged := append(u.Bonuses, bonuses...)
	for _, b := range merged {
		result[b.Type] += b.Count
	}
	r := make([]bonus.Bonus, 0)
	for k, v := range result {
		b := bonus.Bonus{
			Type:  k,
			Count: v,
		}
		r = append(r, b)
	}
	u.Bonuses = r
}

func (u *User) buildLetters(bonuses []bonus.Bonus) {
	for _, bonus := range bonuses {
		if bonus.Type == "LetterBonus" {
			newLetters := getRandomLetterBonuses(bonus.Count)
			u.mergeLetters(newLetters)
			break
		}
	}
}

func getRandomLetterBonuses(count int) []bonus.Bonus {
	letters := make([]bonus.Bonus, count)
	ls := game.NewLetterSet(0, 0)

	for i := 0; i < count; i++ {
		letter := ls.GetRandomLetter()
		letterBonus := bonus.Bonus{
			Value: letter.Value,
			Count: 1,
		}
		letters[i] = letterBonus
	}
	return letters
}

func (u *User) mergeLetters(letters []bonus.Bonus) {
	result := make(map[string]int)
	merged := append(u.Letters, letters...)
	for _, b := range merged {
		if _, ok := result[b.Value]; ok {
			result[b.Value] += b.Count
		} else {
			result[b.Value] = b.Count
		}
	}
	r := make([]bonus.Bonus, 0)
	for k, v := range result {
		b := bonus.Bonus{
			Value: k,
			Count: v,
		}
		r = append(r, b)
	}
	u.Letters = r
}
