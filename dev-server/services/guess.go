package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type GuessState string

const (
	CorrectGuess       GuessState = "CorrectGuess"
	None                          = "None"
	PartlyCorrectGuess            = "PartlyCorrectGuess"
	WrongGuess                    = "WrongGuess"
)

type Pokemons struct {
	Pokemons []Pokemon `json:"pokemons"`
}

type Pokemon struct {
	name           string `json:"name"`
	type1          string `json:"type1"`
	type2          string `json:"type2"`
	generation     int    `json:"generation"`
	evolutionState int    `json:"evolutionState"`
	isFullyEvolved string `json:"isFullyEvolved"`
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

	jsonFile, err := os.Open("./data/pokemonData-v2.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened jsonFile")

	defer jsonFile.Close()

	fmt.Println(*jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var pokemons Pokemons

	json.Unmarshal(byteValue, &pokemons)

	fmt.Println("Slice length: " + strconv.Itoa(len(pokemons.Pokemons)))

	for i := 0; i < 10; i++ {
		fmt.Println("Pokemon name: " + pokemons.Pokemons[i].name)
	}
}
