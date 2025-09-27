package repositories

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoRepository struct {
	Collection *mongo.Collection
}

func NewUserMongoRepository(db *mongo.Database) interfaces.UserRepository {
	return &UserMongoRepository{
		Collection: db.Collection("user_v1"),
	}
}

func (repo *UserMongoRepository) CheckUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	if err := repo.Collection.FindOne(ctx, bson.M{"username": username}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
