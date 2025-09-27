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

type KritikSaranMongoRepository struct {
	Collection *mongo.Collection
}

func NewKritikSaranMongoRepository(db *mongo.Database) interfaces.KritikSaranRepository {
	return &KritikSaranMongoRepository{
		Collection: db.Collection("kritik_saran_v1"),
	}
}

func (repo *KritikSaranMongoRepository) GetKritikSaranFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.KritikSaran, int64, error) {
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

	if tahun != "" {
		filter["tahun"] = tahun // tetap string, sesuai DB
	}

	if semester != "" {
		if semester == "ganjil" {
			filter["semester"] = bson.M{"$in": []string{"1", "3", "5", "7"}}
		} else if semester == "genap" {
			filter["semester"] = bson.M{"$in": []string{"2", "4", "6", "8"}}
		} else {
			// kalau langsung dikirim semester angka (contoh: "7")
			filter["semester"] = semester
		}
	}

	if search != "" {
		filter["$text"] = bson.M{"$search": search}
	}

	var wg sync.WaitGroup
	wg.Add(2)

	var results []models.KritikSaran
	var total int64
	var dataErr, countErr error

	go func() {
		defer wg.Done()
		findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit)).SetSort(bson.D{
			{Key: "tahun", Value: -1},
			{Key: "semester", Value: -1},
			{Key: "_id", Value: 1},
		})

		if search != "" {
			findOptions.SetProjection(bson.M{"score": bson.M{"$meta": "textScore"}})
			findOptions.SetSort(bson.D{
				{Key: "score", Value: bson.M{"$meta": "textScore"}},
				{Key: "tahun", Value: -1},
				{Key: "semester", Value: -1},
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
