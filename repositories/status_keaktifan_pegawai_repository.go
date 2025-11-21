package repositories

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StatusKeaktifanPegawaiMongoRepository struct {
	Collection *mongo.Collection
}

func NewStatusKeaktifanPegawaiMongoRepository(db *mongo.Database) interfaces.StatusKeaktifanPegawaiRepository {
	return &StatusKeaktifanPegawaiMongoRepository{
		Collection: db.Collection("status_keaktifan_pegawai"),
	}
}

func (repo *StatusKeaktifanPegawaiMongoRepository) GetStatusKeaktifanPegawai(ctx context.Context) ([]models.Status, error) {
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
