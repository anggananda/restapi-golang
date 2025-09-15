package repositories

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type PerpemMongoRepository struct {
	Collection *mongo.Collection
}

func NewPerpemMongoRepository(db *mongo.Database) interfaces.PerpemRepository {
	return &PerpemMongoRepository{
		Collection: db.Collection("perpem_v2"),
	}
}

func (repo *PerpemMongoRepository) GetPerpemFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester string) ([]models.Perpem, error) {
  
}
