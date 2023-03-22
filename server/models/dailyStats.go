package models

type DailyStats struct {
	Date                      string `json:"date"`
	ClassicGamesWon           int    `json:"classicGamesWon"`
	ClassicFirstTryWins       int    `json:"classicFirstTryWins"`
	ClassicNumberOfGuesses    int    `json:"classicNumberOfGuesses"`
	FlavortextGamesWon        int    `json:"flavortextGamesWon"`
	FlavortextFirstTryWins    int    `json:"flavortextFirstTryWins"`
	FlavortextNumberOfGuesses int    `json:"flavortextNumberOfGuesses"`
}
