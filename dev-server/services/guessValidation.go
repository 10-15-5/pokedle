package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gabr0236/pokedle/dev-server/models"
)

func SubmitGuess(name string) {
	fmt.Printf("Name passed to SubmitGuess: %v\n", name)

	content, err := os.ReadFile("./data/pokemonData-v2.json")

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	fmt.Println("Successfully Opened jsonFile")

	var pokemons []models.Pokemon

	err = json.Unmarshal(content, &pokemons)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("Pokemon name: %v\n", pokemons[i].Name)
		fmt.Printf("Pokemon name: %v\n", pokemons[i])
	}
}
