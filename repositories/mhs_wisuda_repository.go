package repositories

import (
	"context"
	"fmt"
	"restapi-golang/interfaces"
	"restapi-golang/models"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MhsWisudaMongoRepository struct {
	DB *mongo.Database
}

func NewMhsWisudaMongoRepository(db *mongo.Database) interfaces.MhsWisudaRepository {
	return &MhsWisudaMongoRepository{
		DB: db,
	}
}

func (repo *MhsWisudaMongoRepository) getCollectionByYear(year int) *mongo.Collection {
	return repo.DB.Collection(fmt.Sprintf("mahasiswa_wisuda_%d", year))
}

func (repo *MhsWisudaMongoRepository) GetMhsWisudaFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, search string, tahun, bulan, page, limit int) ([]models.MhsWisuda, int64, error) {
	skip := (page - 1) * limit
filter := bson.M{}

	if kodeFakultas != "" {
		filter["unit.fkt_kode"] = kodeFakultas
	}

	if kodeJurusan != "" {
		filter["unit.jrs_kode"] = kodeJurusan
	}

	if kodeProdi != "" {
		filter["unit.prd_kode"] = kodeProdi
	}

	if bulan != 0 {
		filter["bulan_wisuda"] = bulan
	}

	if search != "" {
		filter["$text"] = bson.M{"$search": search}
	}

	var wg sync.WaitGroup
	wg.Add(2)
	var results []models.MhsWisuda
	var total int64
	var dataErr, countErr error

	go func() {
		defer wg.Done()
		findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit)).SetSort(bson.D{
			{Key: "bulan_wisuda", Value: -1},
			{Key: "_id", Value: 1},
		})

		if search != "" {
			findOptions.SetProjection(bson.M{"score": bson.M{"$meta": "textScore"}})
			findOptions.SetSort(bson.D{
				{Key: "score", Value: bson.M{"$meta": "textScore"}},

				{Key: "bulan_wisuda", Value: -1},
				{Key: "_id", Value: 1},
			})
		}

		cursor, err := repo.getCollectionByYear(tahun).Find(ctx, filter, findOptions)
		if err != nil {
			dataErr = err
			return
		}
		defer cursor.Close(ctx)
		if err := cursor.All(ctx, &results); err != nil {
			dataErr = err
			return
		}
	}()

	go func() {
		defer wg.Done()
		total, countErr = repo.getCollectionByYear(tahun).CountDocuments(ctx, filter)
	}()

	wg.Wait()

	if dataErr != nil {
		return nil, 0, fmt.Errorf("failed to get data: %v", dataErr)
	}
	if countErr != nil {
		return nil, 0, fmt.Errorf("failed to count documents: %v", countErr)
	}

	return results, total, nil

}
