package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GameWon struct {
	CreatedAt       time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	NumberOfGuesses int       `json:"numberOfGuesses,omitempty" bson:"numberOfGuesses,omitempty"`
}

type User struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt" bson:"updatedAt"`
	CurrentStreak int                `json:"currentStreak" bson:"currentStreak"`
	MaxStreak     int                `json:"maxStreak" bson:"maxStreak"`
	FirstTryWins  int                `json:"firstTryWins" bson:"firstTryWins"`
	GamesWon      []GameWon          `json:"gamesWon" bson:"gamesWon"`
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
