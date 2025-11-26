package repositories

import (
	"context"
	"fmt"
	"restapi-golang/interfaces"
	"restapi-golang/models"
	"restapi-golang/utils"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DosenMongoRepository struct {
	Collection *mongo.Collection
}

func NewDosenMongoRepository(db *mongo.Database) interfaces.DosenRepository {
	return &DosenMongoRepository{
		Collection: db.Collection("dosen"),
	}
}

func (repo *DosenMongoRepository) GetDetailDosen(ctx context.Context, niu string) (*models.Dosen, error) {
	var dosen models.Dosen

	if err := repo.Collection.FindOne(ctx, bson.M{"no_induk_undiksha": niu}).Decode(&dosen); err != nil {
		return nil, err
	}

	return &dosen, nil
}

func (repo *DosenMongoRepository) GetDosenHistoryFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, search string, tahun, statusPegawai, statusKeaktifan, page, limit int) ([]models.DosenHistoryResponse, int64, error) {
	normalMatch := bson.M{
		"history.tahun": tahun,
	}

	if kodeFakultas != "" {
		normalMatch["unit.fkt_kode"] = kodeFakultas
	}
	if kodeJurusan != "" {
		normalMatch["unit.jrs_kode"] = kodeJurusan
	}
	if kodeProdi != "" {
		normalMatch["unit.prd_kode"] = kodeProdi
	}
	if statusPegawai != 0 {
		normalMatch["history.id_status_pegawai"] = statusPegawai
	}
	if statusKeaktifan != 0 {
		normalMatch["history.id_status_keaktifan"] = statusKeaktifan
	}

	var textStage bson.A
	if search != "" {
		textStage = bson.A{
			bson.M{"$match": bson.M{
				"$text": bson.M{"$search": search},
			}},
		}
	}

	basePipeline := bson.A{
		bson.M{"$unwind": "$history"},
		bson.M{"$match": normalMatch},
	}

	projection := bson.M{
		"$project": bson.M{
			"nip":                1,
			"no_induk_undiksha":  1,
			"nama":               1,
			"jabatan_fungsional": 1,
			"strata":             1,
			"fakultas":           "$unit.fakultas",
			"jurusan":            "$unit.jurusan",
			"prodi":              "$unit.prodi",
			"tahun":              "$history.tahun",
			"status_pegawai":     "$history.status_pegawai",
			"status_keaktifan":   "$history.status_keaktifan",
		},
	}

	if search != "" {
		projection["$project"].(bson.M)["score"] = bson.M{"$meta": "textScore"}
	}

	sortStage := bson.M{"$sort": bson.M{
		"nim":      1,
		"tahun":    1,
		"semester": 1,
	}}

	if search != "" {
		sortStage = bson.M{"$sort": bson.M{
			"score":    bson.M{"$meta": "textScore"},
			"nim":      1,
			"tahun":    1,
			"semester": 1,
		}}
	}

	var paging bson.A

	if limit > 0 {
		skip := (page - 1) * limit
		paging = bson.A{
			bson.M{"$skip": skip},
			bson.M{"$limit": limit},
		}
	} else {
		paging = bson.A{}
	}

	var wg sync.WaitGroup
	wg.Add(2)

	var results []models.DosenHistoryResponse
	var total int64
	var dataErr, countErr error

	go func() {
		defer wg.Done()

		pipeline := append(textStage, basePipeline...)
		pipeline = append(pipeline, projection, sortStage)
		pipeline = append(pipeline, paging...)

		cursor, err := repo.Collection.Aggregate(ctx, pipeline)
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

		countPipeline := append(textStage, basePipeline...)
		countPipeline = append(countPipeline, bson.M{"$count": "total"})

		cursor, err := repo.Collection.Aggregate(ctx, countPipeline)
		if err != nil {
			countErr = err
			return
		}
		defer cursor.Close(ctx)

		var countResult []bson.M
		if err := cursor.All(ctx, &countResult); err != nil {
			countErr = err
			return
		}

		if len(countResult) > 0 {
			total = utils.ConvertToInt64(countResult[0]["total"])
		}
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
