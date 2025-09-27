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

type RekapPMBMongoRepository struct {
	Collection *mongo.Collection
}

func NewRekapPMBMongoRepository(db *mongo.Database) interfaces.RekapPMBRepository {
	return &RekapPMBMongoRepository{
		Collection: db.Collection("rekap_pmb_v1"),
	}
}

func (repo *RekapPMBMongoRepository) GetRekapPMBFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, search string, tahun, page, limit int) ([]models.RekapPMB, int64, error) {
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

	if tahun != 0 {
		filter["tahun"] = tahun
	}

	if search != "" {
		filter["$text"] = bson.M{"$search": search}
	}

	var wg sync.WaitGroup
	wg.Add(2)
	var results []models.RekapPMB
	var total int64
	var dataErr, countErr error

	go func() {
		defer wg.Done()
		findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit)).SetSort(bson.D{
			{Key: "tahun", Value: -1},
			{Key: "_id", Value: 1},
		})

		if search != "" {
			findOptions.SetProjection(bson.M{"score": bson.M{"$meta": "textScore"}})
			findOptions.SetSort(bson.D{
				{Key: "score", Value: bson.M{"$meta": "textScore"}},
				{Key: "tahun", Value: -1},
				{Key: "_id", Value: 1},
			})
		}

		cursor, err := repo.Collection.Find(ctx, filter, findOptions)
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
		total, countErr = repo.Collection.CountDocuments(ctx, filter)
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
