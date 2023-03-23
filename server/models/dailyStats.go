package models

type DailyStats struct {
	Date                      string `bson:"date" json:"date"`
	ClassicGamesWon           int    `bson:"classicGamesWon" json:"classicGamesWon"`
	ClassicFirstTryWins       int    `bson:"classicFirstTryWins" json:"classicFirstTryWins"`
	ClassicNumberOfGuesses    int    `bson:"classicNumberOfGuesses" json:"classicNumberOfGuesses"`
	FlavortextGamesWon        int    `bson:"flavortextGamesWon" json:"flavortextGamesWon"`
	FlavortextFirstTryWins    int    `bson:"flavortextFirstTryWins" json:"flavortextFirstTryWins"`
	FlavortextNumberOfGuesses int    `bson:"flavortextNumberOfGuesses" json:"flavortextNumberOfGuesses"`
	SilhouetteGamesWon        int    `bson:"silhouetteGamesWon" json:"silhouetteGamesWon"`
	SilhouetteFirstTryWins    int    `bson:"silhouetteFirstTryWins" json:"silhouetteFirstTryWins"`
	SilhouetteNumberOfGuesses int    `bson:"silhouetteNumberOfGuesses" json:"silhouetteNumberOfGuesses"`
}
