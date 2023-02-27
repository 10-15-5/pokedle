package user

import (
	"context"
	"os"
	"time"

	"github.com/gabr0236/pokedle/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collectionName = "users"

// TODO: this is not being used
type UsersRepository interface {
	InsertNewUser(ctx context.Context, user models.User) error
}

type userRepository struct {
	client mongo.Client
}

func GetUserRepository(mongoClient *mongo.Client) *userRepository {
	r := &userRepository{
		client: *mongoClient,
	}
	return r
}

func (r *userRepository) GetUser(userId primitive.ObjectID) models.User {

	coll := r.client.Database(os.Getenv("DATABASE")).Collection(collectionName)

	result := coll.FindOne(context.Background(), bson.M{"_id": userId})

	var user models.User

	result.Decode(&user)

	return user
}

func (r *userRepository) InsertNewUser(user models.User) error {

	coll := r.client.Database(os.Getenv("DATABASE")).Collection(collectionName)

	_, err := coll.InsertOne(context.Background(), user)

	return err
}

func (r *userRepository) InsertNewGameWon(
	userId primitive.ObjectID,
	gameWon models.GameWon,
	currentStreak int,
	maxStreak int,
	isFirstTryWin int,
) models.User {

	coll := r.client.Database(os.Getenv("DATABASE")).Collection(collectionName)

	after := options.After

	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	result := coll.FindOneAndUpdate(context.Background(),
		bson.M{"_id": userId},
		bson.M{
			"$push": bson.M{"gamesWon": gameWon},
			"$set":  bson.M{"currentStreak": currentStreak, "maxStreak": maxStreak, "updatedAt": time.Now()},
			"$inc":  bson.M{"firstTryWins": isFirstTryWin},
		},
		&opt)

	var user models.User

	result.Decode(&user)

	return user
}
