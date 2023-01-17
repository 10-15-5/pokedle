package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type GuessState string

const (
	CorrectGuess       GuessState = "CorrectGuess"
	None                          = "None"
	PartlyCorrectGuess            = "PartlyCorrectGuess"
	WrongGuess                    = "WrongGuess"
)

type Pokemon struct {
	Number         int    `json:"number"`
	Name           string `json:"name"`
	Type1          string `json:"type1"`
	Type2          string `json:"type2"`
	Total          int    `json:"total"`
	Hp             int    `json:"hp"`
	Attack         int    `json:"attack"`
	Defense        int    `json:"defense"`
	SpAtk          int    `json:"sp-atk"`
	SpDef          int    `json:"sp-def"`
	Speed          int    `json:"speed"`
	Generation     int    `json:"generation"`
	Legendary      string `json:"legendary"`
	Evolutions     string `json:"evolutions"`
	EvolutionState int    `json:"evolutionState"`
	IsFullyEvolved string `json:"isFullyEvolved"`
}
type GuessResult struct {
	name           GuessState `json:"name"`
	type1          GuessState `json:"type1"`
	type2          GuessState `json:"type2"`
	generation     GuessState `json:"generation"`
	evolutionState GuessState `json:"evolutionState"`
	isFullyEvolved GuessState `json:"isFullyEvolved"`
}

func SubmitGuess(name string) {
	fmt.Printf("Name passed to SubmitGuess: %v\n", name)

	content, err := os.ReadFile("./data/pokemonData-v2.json")

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	fmt.Println("Successfully Opened jsonFile")

	var pokemons []Pokemon

	err = json.Unmarshal(content, &pokemons)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
}
