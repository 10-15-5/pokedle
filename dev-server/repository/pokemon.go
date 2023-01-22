package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/gabr0236/pokedle/dev-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "Pokemon"

type PokemonRepository interface {
	InsertNewPokemon(ctx context.Context, pokemon models.Pokemon) (bson.M, error)
}

type pokemonRepository struct {
	client mongo.Client
}

func GetRepository(mongoClient *mongo.Client) *pokemonRepository {
	r := &pokemonRepository{
		client: *mongoClient,
	}
	return r
}

func (r *pokemonRepository) InsertNewPokemon(ctx context.Context, pokemon models.Pokemon) (mongo.InsertOneResult, error) {
	fmt.Println("Pokemon:", pokemon)

	coll := r.client.Database(os.Getenv("DATABASE")).Collection(collectionName)
	fmt.Println("Collection fetched")

	result, err := coll.InsertOne(context.TODO(), pokemon)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

	return *result, err
}
