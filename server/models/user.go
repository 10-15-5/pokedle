package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GameWon struct {
	CreatedAt       time.Time `bson:"createdAt"`
	NumberOfGuesses int       `bson:"numberOfGuesses"`
}

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt     time.Time          `bson:"createdAt"`
	UpdatedAt     time.Time          `bson:"updatedAt"`
	CurrentStreak int                `bson:"currentStreak"`
	MaxStreak     int                `bson:"maxStreak"`
	FirstTryWins  int                `bson:"firstTryWins"`
	GamesWon      []GameWon          `bson:"gamesWon"`
}

func NewUser() User {

	objectId := primitive.NewObjectID()

	user := User{
		ID:            objectId,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		CurrentStreak: 0,
		MaxStreak:     0,
		FirstTryWins:  0,
		GamesWon:      []GameWon{},
	}

	return user

}
