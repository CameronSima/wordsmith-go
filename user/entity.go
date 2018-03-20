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
	Username    string             `json:"username"`
	Password    string             `json:"password"`
	Email       string             `json:"email"`
	LevelConfig config.LevelConfig `json:"levelConfig"`
	GameStats   GameStats          `json:"gameStats"`
	IsAdmin     bool               `json:"-"`

	// Letters are stored on the user entity because appengine doesn't allow
	// arrays of arrays ([[]LetterBonus]Bonus)
	Letters []bonus.LetterBonus `json:"letters"`
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
	l := []bonus.LetterBonus{
		bonus.LetterBonus{
			Value: "A",
			Count: 1,
		},
		bonus.LetterBonus{
			Value: "E",
			Count: 2,
		},
		bonus.LetterBonus{
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
	u.mergeBonuses(&userLevel)
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
			if i > 0 {
				config.PointsToNextLevel = l[i-1].PointsRequired
			}
			userLevel = config
			break
		}
	}
	return userLevel
}

// when a user levels up, he gets new bonuses and carries over his
// unused ones.
func (u *User) mergeBonuses(userLevel *config.LevelConfig) {
	for _, earnedBonus := range userLevel.Bonuses {
		for _, ownedBonus := range u.LevelConfig.Bonuses {
			if earnedBonus.Type == ownedBonus.Type {
				earnedBonus.Count += ownedBonus.Count
			}
		}
	}
}

// AddBonus to user
func (u *User) AddBonus(bonusName string, count int) {
	if len(u.LevelConfig.Bonuses) == 0 {
		u.addNewBonus(bonusName, count)
	} else {
		u.incrementBonus(bonusName, count)
	}
}

func (u *User) addNewBonus(bonusName string, count int) {
	newBonuses := make([]bonus.Bonus, 1)
	newBonuses[0] = bonus.Bonus{
		Type:  bonusName,
		Count: count,
	}
	u.LevelConfig.Bonuses = newBonuses
}

func (u *User) incrementBonus(bonusName string, count int) {
	for i := range u.LevelConfig.Bonuses {
		b := &u.LevelConfig.Bonuses[i]
		if b.Type == bonusName {
			b.Count += count
			break
		}
	}
}

// CheckPassword checks a user's saved hashed password against a string.
func (u User) CheckPassword(pw string) error {
	hashedPwBtyeArr := []byte(u.Password)
	pwBtyeArr := []byte(pw)
	return bcrypt.CompareHashAndPassword(hashedPwBtyeArr, pwBtyeArr)
}

func HashPassword(pw string) (string, error) {
	byteArr := []byte(pw)
	hashed, err := bcrypt.GenerateFromPassword(byteArr, bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("Couldn't hash password")
	}
	return string(hashed), nil
}
