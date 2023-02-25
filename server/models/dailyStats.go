package models

type DailyStats struct {
	Date         string `json:"date"`
	GamesWon     int    `json:"gamesWon"`
	FirstTryWins int    `json:"firstTryWins"`
}
