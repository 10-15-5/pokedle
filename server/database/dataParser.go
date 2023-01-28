package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gabr0236/pokedle/server/models"
)

func GetPokemonData() []models.Pokemon {

	content, err := os.ReadFile("./database/pokemonData-v3.json")

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	fmt.Println("Successfully Opened jsonFile")

	var pokemons []models.Pokemon

	err = json.Unmarshal(content, &pokemons)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return pokemons
}
