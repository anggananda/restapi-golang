package repositories

import (
	"context"
	"fmt"
	"restapi-golang/interfaces"
	"restapi-golang/models"
	"restapi-golang/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DashboardPegawaiMongoRepository struct {
	DB *mongo.Database
}

func NewDashboardPegawaiRepository(db *mongo.Database) interfaces.DashboardPegawaiRepository {
	return &DashboardPegawaiMongoRepository{
		DB: db,
	}
}

func (repo *DashboardPegawaiMongoRepository) getCollectionByYear(year int) *mongo.Collection {
	return repo.DB.Collection(fmt.Sprintf("dashboard_pegawai_%d", year))
}

func (repo *DashboardPegawaiMongoRepository) GetDashboardPegawaiOverview(ctx context.Context, tahun int) ([]models.DashboardCardPegawai, error) {

	pipeline := bson.A{
		bson.M{"$group": bson.M{
			"_id": bson.M{
				"id_status_pegawai":   "$id_status_pegawai",
				"id_status_keaktifan": "$id_status_keaktifan",
			},
			"total":            bson.M{"$sum": "$jumlah"},
			"status_pegawai":   bson.M{"$first": "$status_pegawai"},
			"status_keaktifan": bson.M{"$first": "$status_keaktifan"},
		}},
		bson.M{"$sort": bson.M{
			"_id.id_status_pegawai":   1,
			"_id.id_status_keaktifan": 1,
		}},
	}

	cursor, err := repo.getCollectionByYear(tahun).Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	var cards []models.DashboardCardPegawai
	for _, item := range results {
		idDoc := item["_id"].(bson.M)
		statusPegawai := item["status_pegawai"].(string)
		statusKeaktifan := item["status_keaktifan"].(string)
		combinedTitle := fmt.Sprintf("%s (%s)", statusPegawai, statusKeaktifan)
		total, ok := item["total"].(int64)
		if !ok {
			total = utils.ConvertToInt64(item["total"])
		}

		cards = append(cards, models.DashboardCardPegawai{
			Title:             combinedTitle,
			Value:             total,
			IDStatusPegawai:   int64(idDoc["id_status_keaktifan"].(int32)),
			IDStatusKeaktifan: int64(idDoc["id_status_keaktifan"].(int32)),
			Drilldown:         true,
		})
	}

	return cards, nil
}

func (repo *DashboardPegawaiMongoRepository) GetDrilldownPegawaiFakultas(ctx context.Context, tahun, statusPegawai, statusKeaktifan int) ([]models.DrilldownItem, int64, error) {

	matchConditions := bson.M{
		"id_status_pegawai":   statusPegawai,
		"id_status_keaktifan": statusKeaktifan,
	}

	pipeline := bson.A{
		bson.M{"$match": matchConditions},
		bson.M{"$group": bson.M{
			"_id": bson.M{
				"fkt_kode": "$unit.fkt_kode",
				"fakultas": "$unit.fakultas",
			},
			"total": bson.M{"$sum": "$jumlah"},
		}},
		bson.M{"$sort": bson.M{"total": -1}},
	}

	cursor, err := repo.getCollectionByYear(tahun).Aggregate(ctx, pipeline)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, 0, err
	}

	var items []models.DrilldownItem
	var total int64

	for _, item := range results {
		id := item["_id"].(bson.M)
		value := utils.ConvertToInt64(item["total"])

		items = append(items, models.DrilldownItem{
			ID:    id["fkt_kode"].(string),
			Name:  id["fakultas"].(string),
			Value: value,
			Level: "fakultas",
		})
		total += value
	}

	return items, total, nil
}

func (repo *DashboardPegawaiMongoRepository) GetDrilldownPegawaiJurusan(ctx context.Context, tahun, statusPegawai, statusKeaktifan int, kodeFakultas string) ([]models.DrilldownItem, int64, error) {
	matchConditions := bson.M{
		"id_status_pegawai":   statusPegawai,
		"id_status_keaktifan": statusKeaktifan,
		"unit.fkt_kode":       kodeFakultas,
	}

	pipeline := bson.A{
		bson.M{"$match": matchConditions},
		bson.M{"$group": bson.M{

			"_id": bson.M{
				"jrs_kode": "$unit.jrs_kode",
				"jurusan":  "$unit.jurusan",
			},
			"total": bson.M{"$sum": "$jumlah"},
		}},
		bson.M{"$sort": bson.M{"total": -1}},
	}

	cursor, err := repo.getCollectionByYear(tahun).Aggregate(ctx, pipeline)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, 0, err
	}

	var items []models.DrilldownItem
	var total int64

	for _, item := range results {
		id := item["_id"].(bson.M)
		value := utils.ConvertToInt64(item["total"])

		items = append(items, models.DrilldownItem{
			ID:    id["jrs_kode"].(string),
			Name:  id["jurusan"].(string),
			Value: value,
			Level: "jurusan",
		})
		total += value
	}

	return items, total, nil
}

func (repo *DashboardPegawaiMongoRepository) GetDrilldownPegawaiProdi(ctx context.Context, tahun, statusPegawai, statusKeaktifan int, kodeJurusan string) ([]models.DrilldownItem, int64, error) {
	matchConditions := bson.M{
		"id_status_pegawai":   statusPegawai,
		"id_status_keaktifan": statusKeaktifan,
		"unit.jrs_kode":       kodeJurusan,
	}

	pipeline := bson.A{
		bson.M{"$match": matchConditions},
		bson.M{"$group": bson.M{
			"_id": bson.M{
				"prd_kode": "$unit.prd_kode",
				"prodi":    "$unit.prodi",
			},
			"total": bson.M{"$sum": "$jumlah"},
		}},
		bson.M{"$sort": bson.M{"total": -1}},
	}

	cursor, err := repo.getCollectionByYear(tahun).Aggregate(ctx, pipeline)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, 0, err
	}

	var items []models.DrilldownItem
	var total int64

	for _, item := range results {
		id := item["_id"].(bson.M)
		value := utils.ConvertToInt64(item["total"])

		items = append(items, models.DrilldownItem{
			ID:    id["prd_kode"].(string),
			Name:  id["prodi"].(string),
			Value: value,
			Level: "prodi",
		})
		total += value
	}

	return items, total, nil
}
