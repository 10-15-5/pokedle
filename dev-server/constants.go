package main

type GuessState string

const (
	CorrectGuess       GuessState = "CorrectGuess"
	None                          = "None"
	PartlyCorrectGuess            = "PartlyCorrectGuess"
	WrongGuess                    = "WrongGuess"
)
