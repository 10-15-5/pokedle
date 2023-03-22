package secretPokemon

import (
	"context"
	"fmt"
	"os"

	"github.com/gabr0236/pokedle/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: in const file?
const collectionName = "secretPokemons"

// TODO: this is not being used
type PokemonRepository interface {
	InsertNewPokemon(ctx context.Context, pokemon models.Pokemon) (bson.M, error)
}

type pokemonRepository struct {
	client mongo.Client
}

func GetPokemonRepository(mongoClient *mongo.Client) *pokemonRepository {
	r := &pokemonRepository{
		client: *mongoClient,
	}
	return r
}

func (r *pokemonRepository) InsertNewPokemon(ctx context.Context, pokemon models.Pokemon, collection string) (*mongo.InsertOneResult, error) {
	fmt.Println("Pokemon:", pokemon)

	coll := r.client.Database(os.Getenv("DATABASE")).Collection(collection)
	fmt.Println("Collection fetched")

	result, err := coll.InsertOne(context.TODO(), pokemon)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

	return result, err
}

func (r *pokemonRepository) FindCurrentSecretPokemon(ctx context.Context, collection string) (models.Pokemon, error) {
	coll := r.client.Database(os.Getenv("DATABASE")).Collection(collection)

	options := options.FindOne().SetSort(bson.M{"_id": -1})

	var secretPokemon models.Pokemon

	err := coll.FindOne(ctx, bson.D{}, options).Decode(&secretPokemon)

	return secretPokemon, err
}

func (r *pokemonRepository) FindLastTwoSecretPokemon(ctx context.Context, collection string) ([]models.Pokemon, error) {
	coll := r.client.Database(os.Getenv("DATABASE")).Collection(collection)

	options := options.Find().SetSort(bson.M{"_id": -1}).SetLimit(2)

	var secretPokemons []models.Pokemon

	cursor, err := coll.Find(ctx, bson.D{}, options)

	if err := cursor.All(context.Background(), &secretPokemons); err != nil {
		panic(err)
	}

	return secretPokemons, err
}

func (r *pokemonRepository) FindLastSecretPokemons(ctx context.Context, limit int, collection string) []models.Pokemon {

	coll := r.client.Database(os.Getenv("DATABASE")).Collection(collection)

	var secretPokemons []models.Pokemon

	options := options.Find()

	options.SetSort(bson.M{"_id": -1})

	options.SetLimit(int64(limit))

	cursor, _ := coll.Find(context.Background(), bson.M{}, options)

	cursor.All(context.TODO(), &secretPokemons)

	return secretPokemons
}
