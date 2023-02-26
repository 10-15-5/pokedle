package user

import (
	"context"
	"os"

	"github.com/gabr0236/pokedle/server/models"
	"go.mongodb.org/mongo-driver/mongo"
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

func (r *userRepository) InsertNewUser(user models.User) error {

	coll := r.client.Database(os.Getenv("DATABASE")).Collection(collectionName)

	_, err := coll.InsertOne(context.Background(), user)

	return err
}
