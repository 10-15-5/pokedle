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
	ID                      primitive.ObjectID `json:"_id" bson:"_id"`
	CreatedAt               time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt               time.Time          `json:"updatedAt" bson:"updatedAt"`
	ClassicCurrentStreak    int                `json:"classicCurrentStreak" bson:"classicCurrentStreak"`
	ClassicMaxStreak        int                `json:"classicMaxStreak" bson:"classicMaxStreak"`
	ClassicFirstTryWins     int                `json:"classicFirstTryWins" bson:"classicFirstTryWins"`
	ClassicGamesWon         []GameWon          `json:"classicGamesWon" bson:"classicGamesWon"`
	FlavortextCurrentStreak int                `json:"flavortextCurrentStreak" bson:"flavortextCurrentStreak"`
	FlavortextMaxStreak     int                `json:"flavortextMaxStreak" bson:"flavortextMaxStreak"`
	FlavortextFirstTryWins  int                `json:"flavortextFirstTryWins" bson:"flavortextFirstTryWins"`
	FlavortextGamesWon      []GameWon          `json:"flavortextGamesWon" bson:"flavortextGamesWon"`
}

func NewUser() User {

	objectId := primitive.NewObjectID()

	user := User{
		ID:                      objectId,
		CreatedAt:               time.Now(),
		UpdatedAt:               time.Now(),
		ClassicCurrentStreak:    0,
		ClassicMaxStreak:        0,
		ClassicFirstTryWins:     0,
		ClassicGamesWon:         []GameWon{},
		FlavortextCurrentStreak: 0,
		FlavortextMaxStreak:     0,
		FlavortextFirstTryWins:  0,
		FlavortextGamesWon:      []GameWon{},
	}

	return user
}
