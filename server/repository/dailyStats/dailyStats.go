package dailyStats

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: in const file?
const collectionName = "dailyStats"

type dailyStatsRepository struct {
	client mongo.Client
}

func GetPokemonRepository(mongoClient *mongo.Client) *dailyStatsRepository {
	r := &dailyStatsRepository{
		client: *mongoClient,
	}
	return r
}

func (r *dailyStatsRepository) UpdateDailyGuessCount(ctx context.Context, date string) error {
	coll := r.client.Database(os.Getenv("DATABASE")).Collection(collectionName)

	filter := bson.D{{"date", date}}
	update := bson.D{{"$inc", bson.D{{"gamesWon", 1}}}}
	options := options.Update().SetUpsert(true)

	_, err := coll.UpdateOne(context.Background(), filter, update, options)

	return err
}
