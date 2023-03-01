package dailyStats

import (
	"context"
	"os"

	"github.com/gabr0236/pokedle/server/models"
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

func (r *dailyStatsRepository) UpdateDailyGuessCount(ctx context.Context, date string, isFirstTryWin int, numberOfGuesses int) (models.DailyStats, error) {
	coll := r.client.Database(os.Getenv("DATABASE")).Collection(collectionName)

	options := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	var result models.DailyStats
	err := coll.FindOneAndUpdate(
		context.Background(),
		bson.M{"date": date},
		bson.M{
			"$inc": bson.M{
				"gamesWon":        1,
				"firstTryWins":    isFirstTryWin,
				"numberOfGuesses": numberOfGuesses,
			},
		},
		options,
	).Decode(&result)

	return result, err
}

func (r *dailyStatsRepository) GetDailyStats(ctx context.Context, date string) (models.DailyStats, error) {
	coll := r.client.Database(os.Getenv("DATABASE")).Collection(collectionName)

	var result models.DailyStats
	err := coll.FindOne(context.Background(), bson.M{"date": date}).Decode(&result)

	return result, err
}
