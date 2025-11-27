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
		Collection: db.Collection("user_auth"),
	}
}

func (repo *UserMongoRepository) CheckUserByUsername(ctx context.Context, username string) (*models.UserAuth, error) {
	var user models.UserAuth
	if err := repo.Collection.FindOne(ctx, bson.M{"auth_info.email_sso": username}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
