package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gabr0236/pokedle/dev-server/database"
	"github.com/gabr0236/pokedle/dev-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SubmitGuess(name string) {
	fmt.Printf("Name passed to SubmitGuess: %v\n", name)

	content, err := os.ReadFile("./database/pokemonData-v2.json")

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	fmt.Println("Successfully Opened jsonFile")

	var pokemons []models.Pokemon

	err = json.Unmarshal(content, &pokemons)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	mongoClient := database.StartMongoDB()

	collection := mongoClient.Database("Dev").Collection("pokemons")

	var result bson.M

	err = collection.FindOne(context.TODO(), bson.D{{"name", "Charizard"}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the name %s\n", "Charizard")
		return
	}
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("Pokemon name: %v\n", pokemons[i].Name)
	// 	fmt.Printf("Pokemon name: %v\n", pokemons[i])
	// }
}
