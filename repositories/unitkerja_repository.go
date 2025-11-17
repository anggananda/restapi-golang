package repositories

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UnitKerjaMongoRepository struct {
	Collection *mongo.Collection
}

func NewUnitKerjaMongoRepository(db *mongo.Database) interfaces.UnitKerjaRepository {
	return &UnitKerjaMongoRepository{
		Collection: db.Collection("unitkerja"),
	}
}

func (repo *UnitKerjaMongoRepository) GetUnitKerja(ctx context.Context) (*models.UnitKerjaMapping, error) {
	var unitkerja models.UnitKerjaMapping
	if err := repo.Collection.FindOne(ctx, bson.M{"_id": "unit_kerja"}).Decode(&unitkerja); err != nil {
		return nil, err
	}
	return &unitkerja, nil
}
