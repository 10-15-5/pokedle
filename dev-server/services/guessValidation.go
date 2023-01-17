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
	Name           string `json:"name"`
	Type1          string `json:"type1"`
	Type2          string `json:"type2"`
	Generation     int    `json:"generation"`
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

	for i := 0; i < 10; i++ {
		fmt.Printf("Pokemon name: %v\n", pokemons[i].Name)
		fmt.Printf("Pokemon name: %v\n", pokemons[i])
	}
}
