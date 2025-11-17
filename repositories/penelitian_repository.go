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

type PenelitianMongoRepository struct {
	DB *mongo.Database
}

func NewPenelitianMongoRepository(db *mongo.Database) interfaces.PenelitianRepository {
	return &PenelitianMongoRepository{
		DB: db,
	}
}

func (repo *PenelitianMongoRepository) getCollectionByYear(year string) *mongo.Collection {
	return repo.DB.Collection(fmt.Sprintf("penelitian_%s", year))
}

func (repo *PenelitianMongoRepository) GetPenelitianFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.Penelitian, int64, error) {
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

	if semester != "" {
		if semester == "ganjil" {
			filter["cron_semester"] = "1"
		} else if semester == "genap" {
			filter["cron_semester"] = "2"
		} else {
			filter["cron_semester"] = semester
		}
	}

	if search != "" {
		filter["$text"] = bson.M{"$search": search}
	}

	var wg sync.WaitGroup
	wg.Add(2)

	var results []models.Penelitian
	var total int64
	var dataErr, countErr error

	go func() {
		defer wg.Done()
		findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit)).SetSort(bson.D{
			{Key: "cron_semester", Value: -1},
			{Key: "_id", Value: 1},
		})

		if search != "" {
			findOptions.SetProjection(bson.M{"score": bson.M{"$meta": "textScore"}})
			findOptions.SetSort(bson.D{
				{Key: "score", Value: bson.M{"$meta": "textScore"}},
				{Key: "cron_semester", Value: -1},
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
