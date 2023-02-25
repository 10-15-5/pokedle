package models

import (
	"time"
)

type GameWon struct {
	CreatedAt       time.Time `bson:"createdAt"`
	NumberOfGuesses int       `bson:"numberOfGuesses"`
}

type User struct {
	UserId        string    `bson:"userId"`
	CreatedAt     time.Time `bson:"createdAt"`
	UpdatedAt     time.Time `bson:"updatedAt"`
	CurrentStreak int       `bson:"currentStreak"`
	MaxStreak     int       `bson:"maxStreak"`
	FirstTryWins  int       `bson:"firstTryWins"`
	GamesWon      []GameWon `bson:"gamesWon"`
}
