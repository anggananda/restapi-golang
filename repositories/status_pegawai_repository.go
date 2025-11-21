package repositories

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StatusPegawaiMongoRepository struct {
	Collection *mongo.Collection
}

func NewStatusPegawaiMongoRepository(db *mongo.Database) interfaces.StatusPegawaiRepository {
	return &StatusPegawaiMongoRepository{
		Collection: db.Collection("status_pegawai"),
	}
}

func (repo *StatusPegawaiMongoRepository) GetStatusPegawai(ctx context.Context) ([]models.Status, error) {
	cursor, err := repo.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var statusMhs []models.Status
	if err := cursor.All(ctx, &statusMhs); err != nil {
		return nil, err
	}

	return statusMhs, nil
}
