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

type DashboardMhsRepository struct {
	DB *mongo.Database
}

func NewDashboardMhsRepository(db *mongo.Database) interfaces.DashboardMhsRepository {
	return &DashboardMhsRepository{
		DB: db,
	}
}

func (repo *DashboardMhsRepository) getCollectionByYear(year int) *mongo.Collection {
	return repo.DB.Collection(fmt.Sprintf("dashboard_mahasiswa_%d", year))
}

func (repo *DashboardMhsRepository) HasJurusan(ctx context.Context, tahun int, semester int, status string, kodeFakultas string) (bool, error) {
	filter := bson.M{
		"semester":      semester,
		"id_status":     status,
		"unit.fkt_kode": kodeFakultas,
		"unit.jrs_kode": bson.M{"$ne": nil},
	}

	count, err := repo.getCollectionByYear(tahun).CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repo *DashboardMhsRepository) GetDashboardMhsOverview(ctx context.Context, tahun int, semester int) ([]models.DashboardCard, error) {

	pipeline := bson.A{
		bson.M{"$match": bson.M{
			"semester": semester,
		}},
		bson.M{"$group": bson.M{
			"_id":    "$id_status",
			"total":  bson.M{"$sum": "$jumlah"},
			"status": bson.M{"$first": "$status"},
		}},
		bson.M{"$sort": bson.M{"_id": 1}},
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

	var cards []models.DashboardCard
	for _, item := range results {
		status := item["_id"].(string)
		total, ok := item["total"].(int64)
		if !ok {
			total = utils.ConvertToInt64(item["total"])
		}

		cards = append(cards, models.DashboardCard{
			Title:     item["status"].(string),
			Value:     total,
			Status:    status,
			Drilldown: true,
		})
	}

	return cards, nil
}

func (repo *DashboardMhsRepository) GetDrilldownMhsFakultas(ctx context.Context, tahun int, semester int, status string) ([]models.DrilldownItem, int64, error) {

	matchConditions := bson.M{
		"semester":  semester,
		"id_status": status,
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
		id, ok := item["_id"].(bson.M)
		if !ok {
			continue
		}
		value := utils.ConvertToInt64(item["total"])
		var fktKode string
		if val, ok := id["fkt_kode"].(string); ok {
			fktKode = val
		}
		var namaFakultas string
		if val, ok := id["fakultas"].(string); ok {
			namaFakultas = val
		}
		items = append(items, models.DrilldownItem{
			ID:    fktKode,
			Name:  namaFakultas,
			Value: value,
			Level: "fakultas",
		})
		total += value
	}

	return items, total, nil
}

func (repo *DashboardMhsRepository) GetDrilldownMhsJurusan(ctx context.Context, tahun int, semester int, status string, kodeFakultas string) ([]models.DrilldownItem, int64, error) {
	matchConditions := bson.M{
		"semester":      semester,
		"id_status":     status,
		"unit.fkt_kode": kodeFakultas,
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

		jrsKode, ok := id["jrs_kode"].(string)
		if !ok {
			jrsKode = ""
		}
		jrsName, ok := id["jurusan"].(string)
		if !ok {
			jrsName = ""
		}

		items = append(items, models.DrilldownItem{
			ID:    jrsKode,
			Name:  jrsName,
			Value: value,
			Level: "jurusan",
		})
		total += value
	}

	return items, total, nil
}

func (repo *DashboardMhsRepository) GetDrilldownMhsProdi(ctx context.Context, tahun int, semester int, status string, kodeJurusan string) ([]models.DrilldownItem, int64, error) {
	matchConditions := bson.M{
		"semester":      semester,
		"id_status":     status,
		"unit.jrs_kode": kodeJurusan,
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

func (repo *DashboardMhsRepository) GetProdiByFakultas(
	ctx context.Context,
	tahun int,
	semester int,
	status string,
	kodeFakultas string,
) ([]models.DrilldownItem, int64, error) {

	matchConditions := bson.M{
		"semester":      semester,
		"id_status":     status,
		"unit.fkt_kode": kodeFakultas,
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

		prdKode, _ := id["prd_kode"].(string)
		prodi, _ := id["prodi"].(string)

		items = append(items, models.DrilldownItem{
			ID:    prdKode,
			Name:  prodi,
			Value: value,
			Level: "prodi",
		})
		total += value
	}

	return items, total, nil
}
