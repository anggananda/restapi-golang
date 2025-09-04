// package repositories

// import (
// 	"context"
// 	"fmt"
// 	"restapi-golang/interfaces"
// 	"restapi-golang/models"
// 	"restapi-golang/utils"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type DashboardMhsRepository struct {
// 	Collection *mongo.Collection
// }

// func NewDashboardMhsRepository(db *mongo.Database) interfaces.DashboardMhsRepository {
// 	return &DashboardMhsRepository{
// 		Collection: db.Collection("dashboard_aggregation"),
// 	}
// }

// func (repo *DashboardMhsRepository) GetDashboardOverview(ctx context.Context, tahun int, semester int) ([]models.DashboardCard, error) {
// 	pipeline := bson.A{
// 		bson.M{"$match": bson.M{
// 			"tahun":    tahun,
// 			"semester": semester,
// 		}},
// 		bson.M{"$group": bson.M{
// 			"_id":    "$status_singkatan",
// 			"total":  bson.M{"$sum": "$jumlah"},
// 			"status": bson.M{"$first": "$status"},
// 		}},
// 		bson.M{"$sort": bson.M{"_id": 1}},
// 	}

// 	cursor, err := repo.Collection.Aggregate(ctx, pipeline)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	var results []bson.M
// 	if err = cursor.All(ctx, &results); err != nil {
// 		return nil, err
// 	}

// 	var cards []models.DashboardCard
// 	for _, item := range results {
// 		status := item["_id"].(string)
// 		total := utils.ConvertToInt64(item["total"])

// 		cards = append(cards, models.DashboardCard{
// 			Title:     getStatusTitle(status),
// 			Value:     total,
// 			Icon:      getStatusIcon(status),
// 			Color:     getStatusColor(status),
// 			Status:    status,
// 			Drilldown: true,
// 		})
// 	}

// 	return cards, nil
// }

// func (repo *DashboardMhsRepository) GetDrilldownFakultas(ctx context.Context, tahun int, semester int, status string) ([]models.DrilldownItem, int64, error) {
// 	matchConditions := bson.M{
// 		"tahun":    tahun,
// 		"semester": semester,
// 		"status":   status,
// 	}

// 	pipeline := bson.A{
// 		bson.M{"$match": matchConditions},
// 		bson.M{"$group": bson.M{
// 			"_id": bson.M{
// 				"fakultas": "$fakultas",
// 				"uk_kode":  "$uk_kode",
// 			},
// 			"total": bson.M{"$sum": "$jumlah"},
// 		}},
// 		bson.M{"$sort": bson.M{"total": -1}},
// 	}

// 	cursor, err := repo.Collection.Aggregate(ctx, pipeline)
// 	if err != nil {
// 		return nil, 0, err
// 	}
// 	defer cursor.Close(ctx)

// 	var results []bson.M
// 	if err = cursor.All(ctx, &results); err != nil {
// 		return nil, 0, err
// 	}

// 	var items []models.DrilldownItem
// 	var total int64

// 	for _, item := range results {
// 		id := item["_id"].(bson.M)
// 		value := utils.ConvertToInt64(item["total"])

// 		items = append(items, models.DrilldownItem{
// 			ID:    id["uk_kode"].(string),
// 			Name:  id["fakultas"].(string),
// 			Value: value,
// 			Level: "fakultas",
// 		})
// 		total += value
// 	}

// 	return items, total, nil
// }

// func (repo *DashboardMhsRepository) GetDrilldownJurusan(ctx context.Context, tahun int, semester int, status string, fakultasKode string) ([]models.DrilldownItem, int64, error) {
// 	matchConditions := bson.M{
// 		"tahun":    tahun,
// 		"semester": semester,
// 		"status":   status,
// 		"uk_kode":  bson.M{"$regex": fmt.Sprintf("^%s", fakultasKode)},
// 	}

// 	pipeline := bson.A{
// 		bson.M{"$match": matchConditions},
// 		bson.M{"$group": bson.M{
// 			"_id": bson.M{
// 				"jurusan": "$jurusan",
// 				"uk_kode": "$uk_kode",
// 			},
// 			"total": bson.M{"$sum": "$jumlah"},
// 		}},
// 		bson.M{"$sort": bson.M{"total": -1}},
// 	}

// 	cursor, err := repo.Collection.Aggregate(ctx, pipeline)
// 	if err != nil {
// 		return nil, 0, err
// 	}
// 	defer cursor.Close(ctx)

// 	var results []bson.M
// 	if err = cursor.All(ctx, &results); err != nil {
// 		return nil, 0, err
// 	}

// 	var items []models.DrilldownItem
// 	var total int64

// 	for _, item := range results {
// 		id := item["_id"].(bson.M)
// 		value := utils.ConvertToInt64(item["total"])

// 		items = append(items, models.DrilldownItem{
// 			ID:    id["uk_kode"].(string),
// 			Name:  id["jurusan"].(string),
// 			Value: value,
// 			Level: "jurusan",
// 		})
// 		total += value
// 	}

// 	return items, total, nil
// }

// func (repo *DashboardMhsRepository) GetDrilldownProdi(ctx context.Context, tahun int, semester int, status string, jurusanKode string) ([]models.DrilldownItem, int64, error) {
// 	matchConditions := bson.M{
// 		"tahun":    tahun,
// 		"semester": semester,
// 		"status":   status,
// 		"uk_kode":  bson.M{"$regex": fmt.Sprintf("^%s", jurusanKode)},
// 	}

// 	pipeline := bson.A{
// 		bson.M{"$match": matchConditions},
// 		bson.M{"$group": bson.M{
// 			"_id": bson.M{
// 				"prodi":   "$prodi",
// 				"uk_kode": "$uk_kode",
// 			},
// 			"total": bson.M{"$sum": "$jumlah"},
// 		}},
// 		bson.M{"$sort": bson.M{"total": -1}},
// 	}

// 	cursor, err := repo.Collection.Aggregate(ctx, pipeline)
// 	if err != nil {
// 		return nil, 0, err
// 	}
// 	defer cursor.Close(ctx)

// 	var results []bson.M
// 	if err = cursor.All(ctx, &results); err != nil {
// 		return nil, 0, err
// 	}

// 	var items []models.DrilldownItem
// 	var total int64

// 	for _, item := range results {
// 		id := item["_id"].(bson.M)
// 		value := utils.ConvertToInt64(item["total"])

// 		items = append(items, models.DrilldownItem{
// 			ID:    id["uk_kode"].(string),
// 			Name:  id["prodi"].(string),
// 			Value: value,
// 			Level: "prodi",
// 		})
// 		total += value
// 	}

// 	return items, total, nil
// }

// // Helper functions
// func getStatusTitle(status string) string {
// 	switch status {
// 	case "A":
// 		return "Mahasiswa Aktif"
// 	case "C":
// 		return "Mahasiswa Cuti"
// 	case "D":
// 		return "Drop Out"
// 	default:
// 		return status
// 	}
// }

// func getStatusIcon(status string) string {
// 	switch status {
// 	case "A":
// 		return "👨‍🎓"
// 	case "C":
// 		return "⏸️"
// 	case "D":
// 		return "❌"
// 	default:
// 		return "📊"
// 	}
// }

// func getStatusColor(status string) string {
// 	switch status {
// 	case "A":
// 		return "green"
// 	case "C":
// 		return "orange"
// 	case "D":
// 		return "red"
// 	default:
// 		return "blue"
// 	}
// }

package repositories

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
	"restapi-golang/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DashboardMhsRepository struct {
	Collection *mongo.Collection
}

func NewDashboardMhsRepository(db *mongo.Database) interfaces.DashboardMhsRepository {
	return &DashboardMhsRepository{
		Collection: db.Collection("dashboard_aggre"),
	}
}

func (repo *DashboardMhsRepository) GetDashboardOverview(ctx context.Context, tahun int, semesterType string) ([]models.DashboardCard, error) {
	pipeline := bson.A{
		bson.M{"$match": bson.M{
			"tahun":         tahun,
			"semester_type": semesterType, // Gunakan semester_type bukan semester number
		}},
		bson.M{"$group": bson.M{
			"_id":    "$status_singkatan",
			"total":  bson.M{"$sum": "$jumlah"},
			"status": bson.M{"$first": "$status"},
		}},
		bson.M{"$sort": bson.M{"_id": 1}},
	}

	cursor, err := repo.Collection.Aggregate(ctx, pipeline)
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
		total := utils.ConvertToInt64(item["total"])

		cards = append(cards, models.DashboardCard{
			Title:     getStatusTitle(status),
			Value:     total,
			Icon:      getStatusIcon(status),
			Color:     getStatusColor(status),
			Status:    status,
			Drilldown: true,
		})
	}

	return cards, nil
}

func (repo *DashboardMhsRepository) GetDrilldownFakultas(ctx context.Context, tahun int, semesterType string, status string) ([]models.DrilldownItem, int64, error) {
	matchConditions := bson.M{
		"tahun":         tahun,
		"semester_type": semesterType, // Gunakan semester_type
		"status":        status,       // Perbaiki: gunakan status bukan status
	}

	pipeline := bson.A{
		bson.M{"$match": matchConditions},
		bson.M{"$group": bson.M{
			"_id": bson.M{
				"fakultas": "$fakultas",
				"uk_kode":  "$uk_kode",
			},
			"total": bson.M{"$sum": "$jumlah"},
		}},
		bson.M{"$sort": bson.M{"total": -1}},
	}

	cursor, err := repo.Collection.Aggregate(ctx, pipeline)
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

		// Safe type assertion
		ukKode, ok1 := id["uk_kode"].(string)
		fakultas, ok2 := id["fakultas"].(string)

		if ok1 && ok2 {
			items = append(items, models.DrilldownItem{
				ID:    ukKode,
				Name:  fakultas,
				Value: value,
				Level: "fakultas",
			})
			total += value
		}
	}

	return items, total, nil
}

func (repo *DashboardMhsRepository) GetDrilldownJurusan(ctx context.Context, tahun int, semesterType string, status string, fakultasKode string) ([]models.DrilldownItem, int64, error) {
	matchConditions := bson.M{
		"tahun":         tahun,
		"semester_type": semesterType,
		"status":        status,
		"uk_kode":       fakultasKode, // Exact match, bukan regex
	}

	pipeline := bson.A{
		bson.M{"$match": matchConditions},
		bson.M{"$group": bson.M{
			"_id": bson.M{
				"jurusan": "$jurusan",
				"uk_kode": "$uk_kode",
			},
			"total": bson.M{"$sum": "$jumlah"},
		}},
		bson.M{"$sort": bson.M{"total": -1}},
	}

	cursor, err := repo.Collection.Aggregate(ctx, pipeline)
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

		// Safe type assertion
		ukKode, ok1 := id["uk_kode"].(string)
		jurusan, ok2 := id["jurusan"].(string)

		if ok1 && ok2 {
			items = append(items, models.DrilldownItem{
				ID:    ukKode,
				Name:  jurusan,
				Value: value,
				Level: "jurusan",
			})
			total += value
		}
	}

	return items, total, nil
}

func (repo *DashboardMhsRepository) GetDrilldownProdi(ctx context.Context, tahun int, semesterType string, status string, jurusanKode string) ([]models.DrilldownItem, int64, error) {
	matchConditions := bson.M{
		"tahun":         tahun,
		"semester_type": semesterType,
		"status":        status,
		"uk_kode":       jurusanKode, // Exact match, bukan regex
	}

	pipeline := bson.A{
		bson.M{"$match": matchConditions},
		bson.M{"$group": bson.M{
			"_id": bson.M{
				"prodi":   "$prodi",
				"uk_kode": "$uk_kode",
			},
			"total": bson.M{"$sum": "$jumlah"},
		}},
		bson.M{"$sort": bson.M{"total": -1}},
	}

	cursor, err := repo.Collection.Aggregate(ctx, pipeline)
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

		// Safe type assertion
		ukKode, ok1 := id["uk_kode"].(string)
		prodi, ok2 := id["prodi"].(string)

		if ok1 && ok2 {
			items = append(items, models.DrilldownItem{
				ID:    ukKode,
				Name:  prodi,
				Value: value,
				Level: "prodi",
			})
			total += value
		}
	}

	return items, total, nil
}

// Helper functions (tetap sama)
func getStatusTitle(status string) string {
	switch status {
	case "A":
		return "Mahasiswa Aktif"
	case "C":
		return "Mahasiswa Cuti"
	case "D":
		return "Drop Out"
	default:
		return status
	}
}

func getStatusIcon(status string) string {
	switch status {
	case "A":
		return "👨‍🎓"
	case "C":
		return "⏸️"
	case "D":
		return "❌"
	default:
		return "📊"
	}
}

func getStatusColor(status string) string {
	switch status {
	case "A":
		return "green"
	case "C":
		return "orange"
	case "D":
		return "red"
	default:
		return "blue"
	}
}
