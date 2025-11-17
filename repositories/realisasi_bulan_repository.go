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

type RealisasiBulanMongoRepository struct {
	DB *mongo.Database
}

func NewRealisasiBulanMongoRepository(db *mongo.Database) interfaces.RealisasiBulanRepository {
	return &RealisasiBulanMongoRepository{
		DB: db,
	}
}

func (repo *RealisasiBulanMongoRepository) getCollectionByYear(year string) *mongo.Collection {
	return repo.DB.Collection(fmt.Sprintf("realisasi_bulan_%s", year))
}

func (repo *RealisasiBulanMongoRepository) GetRealisasiBulanFiltered(ctx context.Context, tahun, search string, page, limit int) ([]models.RealisasiBulan, int64, error) {
	skip := (page - 1) * limit
	filter := bson.M{}

	if search != "" {
		filter["$text"] = bson.M{"$search": search}
	}

	var wg sync.WaitGroup
	wg.Add(2)

	var results []models.RealisasiBulan
	var total int64
	var dataErr, countErr error

	go func() {
		defer wg.Done()
		findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit))
		if search != "" {
			findOptions.SetProjection(bson.M{"score": bson.M{"$meta": "textScore"}})
			findOptions.SetSort(bson.D{
				{Key: "score", Value: bson.M{"$meta": "textScore"}},
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
