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
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt     time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt     time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	CurrentStreak int                `json:"currentStreak,omitempty" bson:"currentStreak,omitempty"`
	MaxStreak     int                `json:"maxStreak,omitempty" bson:"maxStreak,omitempty"`
	FirstTryWins  int                `json:"firstTryWins,omitempty" bson:"firstTryWins,omitempty"`
	GamesWon      []GameWon          `json:"gamesWon,omitempty" bson:"gamesWon,omitempty"`
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
