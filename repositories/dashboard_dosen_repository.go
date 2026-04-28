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

type DashboardDosenMongoRepository struct {
	DB *mongo.Database
}

func NewDashboardDosenRepository(db *mongo.Database) interfaces.DashboardDosenRepository {
	return &DashboardDosenMongoRepository{
		DB: db,
	}
}

func (repo *DashboardDosenMongoRepository) getCollectionByYear(year int) *mongo.Collection {
	return repo.DB.Collection(fmt.Sprintf("dashboard_dosen_%d", year))
}

func (repo *DashboardDosenMongoRepository) GetDashboardDosenOverview(ctx context.Context, tahun int) ([]models.DashboardCardPegawai, error) {

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

		combinedTitle := fmt.Sprintf("%s - %s", statusPegawai, statusKeaktifan)

		idPegawai := utils.ConvertToInt64(idDoc["id_status_pegawai"])
		idKeaktifan := utils.ConvertToInt64(idDoc["id_status_keaktifan"])

		total := utils.ConvertToInt64(item["total"])

		cards = append(cards, models.DashboardCardPegawai{
			Title:             combinedTitle,
			Value:             total,
			IDStatusPegawai:   idPegawai,
			IDStatusKeaktifan: idKeaktifan,
			Drilldown:         true,
		})
	}

	return cards, nil
}

func (repo *DashboardDosenMongoRepository) GetDrilldownDosenFakultas(ctx context.Context, tahun, statusPegawai, statusKeaktifan int) ([]models.DrilldownItem, int64, error) {

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

func (repo *DashboardDosenMongoRepository) GetDrilldownDosenJurusan(ctx context.Context, tahun, statusPegawai, statusKeaktifan int, kodeFakultas string) ([]models.DrilldownItem, int64, error) {
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

func (repo *DashboardDosenMongoRepository) GetDrilldownDosenProdi(ctx context.Context, tahun, statusPegawai, statusKeaktifan int, kodeJurusan string) ([]models.DrilldownItem, int64, error) {
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

// type DashboardDosenMongoRepository struct {
// 	DB *mongo.Database
// }

// func NewDashboardDosenRepository(db *mongo.Database) interfaces.DashboardDosenRepository {
// 	return &DashboardDosenMongoRepository{
// 		DB: db,
// 	}
// }

// func (repo *DashboardDosenMongoRepository) getCollectionByYear(year int) *mongo.Collection {
// 	return repo.DB.Collection(fmt.Sprintf("dashboard_dosen_v2_%d", year))
// }

// // ------------------------------------------------------------
// // OVERVIEW
// // Ambil semua metrics dari seluruh prodi (level terdalam),
// // lalu group by status_pegawai + status_keaktifan.
// //
// // Pipeline unwind:
// //
// //	data (universitas)
// //	  → data.children (fakultas)
// //	    → data.children.children (jurusan)
// //	      → data.children.children.children (prodi)
// //	        → data.children.children.children.metrics
// //
// // ------------------------------------------------------------
// func (repo *DashboardDosenMongoRepository) GetDashboardDosenOverview(
// 	ctx context.Context, tahun int,
// ) ([]models.DashboardCardPegawai, error) {

// 	pipeline := bson.A{
// 		// 1. Ambil dokumen tree tahun ini
// 		bson.M{"$match": bson.M{"tahun": tahun}},

// 		// 2. Unwind sampai ke metrics level prodi
// 		bson.M{"$unwind": "$data"},
// 		bson.M{"$unwind": "$data.children"},
// 		bson.M{"$unwind": "$data.children.children"},
// 		bson.M{"$unwind": "$data.children.children.children"},
// 		bson.M{"$unwind": "$data.children.children.children.metrics"},

// 		// 3. Aggregate semua metrics
// 		bson.M{"$group": bson.M{
// 			"_id": bson.M{
// 				"id_status_pegawai":   "$data.children.children.children.metrics.id_status_pegawai",
// 				"id_status_keaktifan": "$data.children.children.children.metrics.id_status_keaktifan",
// 			},
// 			"status_pegawai":   bson.M{"$first": "$data.children.children.children.metrics.status_pegawai"},
// 			"status_keaktifan": bson.M{"$first": "$data.children.children.children.metrics.status_keaktifan"},
// 			"total":            bson.M{"$sum": "$data.children.children.children.metrics.jumlah"},
// 		}},
// 		bson.M{"$sort": bson.M{
// 			"_id.id_status_pegawai":   1,
// 			"_id.id_status_keaktifan": 1,
// 		}},
// 	}

// 	cursor, err := repo.getCollectionByYear(tahun).Aggregate(ctx, pipeline)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	var results []bson.M
// 	if err = cursor.All(ctx, &results); err != nil {
// 		return nil, err
// 	}

// 	var cards []models.DashboardCardPegawai
// 	for _, item := range results {
// 		idDoc := item["_id"].(bson.M)

// 		statusPegawai := item["status_pegawai"].(string)
// 		statusKeaktifan := item["status_keaktifan"].(string)
// 		combinedTitle := fmt.Sprintf("%s - %s", statusPegawai, statusKeaktifan)

// 		idPegawai := utils.ConvertToInt64(idDoc["id_status_pegawai"])
// 		idKeaktifan := utils.ConvertToInt64(idDoc["id_status_keaktifan"])
// 		total := utils.ConvertToInt64(item["total"])

// 		cards = append(cards, models.DashboardCardPegawai{
// 			Title:             combinedTitle,
// 			Value:             total,
// 			IDStatusPegawai:   idPegawai,
// 			IDStatusKeaktifan: idKeaktifan,
// 			Drilldown:         true,
// 		})
// 	}

// 	return cards, nil
// }

// // ------------------------------------------------------------
// // DRILLDOWN FAKULTAS
// // Filter metrics berdasarkan status, lalu group per fakultas
// // (level children ke-1 dari root).
// // ------------------------------------------------------------
// func (repo *DashboardDosenMongoRepository) GetDrilldownDosenFakultas(
// 	ctx context.Context, tahun, statusPegawai, statusKeaktifan int,
// ) ([]models.DrilldownItem, int64, error) {

// 	pipeline := bson.A{
// 		bson.M{"$match": bson.M{"tahun": tahun}},

// 		// Unwind ke level fakultas
// 		bson.M{"$unwind": "$data"},
// 		bson.M{"$unwind": "$data.children"},

// 		// Simpan info fakultas sebelum unwind lebih dalam
// 		bson.M{"$addFields": bson.M{
// 			"fkt_kode": "$data.children.uk_kode",
// 			"fkt_nama": "$data.children.uk_nama",
// 		}},

// 		// Lanjut unwind ke jurusan → prodi → metrics
// 		bson.M{"$unwind": "$data.children.children"},
// 		bson.M{"$unwind": "$data.children.children.children"},
// 		bson.M{"$unwind": "$data.children.children.children.metrics"},

// 		// Filter berdasarkan status yang dipilih
// 		bson.M{"$match": bson.M{
// 			"data.children.children.children.metrics.id_status_pegawai":   statusPegawai,
// 			"data.children.children.children.metrics.id_status_keaktifan": statusKeaktifan,
// 		}},

// 		// Group per fakultas
// 		bson.M{"$group": bson.M{
// 			"_id": bson.M{
// 				"fkt_kode": "$fkt_kode",
// 				"fkt_nama": "$fkt_nama",
// 			},
// 			"total": bson.M{"$sum": "$data.children.children.children.metrics.jumlah"},
// 		}},
// 		bson.M{"$sort": bson.M{"total": -1}},
// 	}

// 	cursor, err := repo.getCollectionByYear(tahun).Aggregate(ctx, pipeline)
// 	if err != nil {
// 		return nil, 0, err
// 	}
// 	defer cursor.Close(ctx)

// 	var results []bson.M
// 	if err = cursor.All(ctx, &results); err != nil {
// 		return nil, 0, err
// 	}

// 	var items []models.DrilldownItem
// 	var grandTotal int64

// 	for _, item := range results {
// 		id := item["_id"].(bson.M)
// 		value := utils.ConvertToInt64(item["total"])

// 		items = append(items, models.DrilldownItem{
// 			ID:    id["fkt_kode"].(string),
// 			Name:  id["fkt_nama"].(string),
// 			Value: value,
// 			Level: "fakultas",
// 		})
// 		grandTotal += value
// 	}

// 	return items, grandTotal, nil
// }

// // ------------------------------------------------------------
// // DRILLDOWN JURUSAN
// // Filter berdasarkan status + fakultas, group per jurusan.
// // ------------------------------------------------------------
// func (repo *DashboardDosenMongoRepository) GetDrilldownDosenJurusan(
// 	ctx context.Context, tahun, statusPegawai, statusKeaktifan int, kodeFakultas string,
// ) ([]models.DrilldownItem, int64, error) {

// 	pipeline := bson.A{
// 		bson.M{"$match": bson.M{"tahun": tahun}},

// 		bson.M{"$unwind": "$data"},
// 		bson.M{"$unwind": "$data.children"},

// 		// Filter hanya fakultas yang dipilih
// 		bson.M{"$match": bson.M{
// 			"data.children.uk_kode": kodeFakultas,
// 		}},

// 		// Simpan info jurusan sebelum unwind lebih dalam
// 		bson.M{"$unwind": "$data.children.children"},
// 		bson.M{"$addFields": bson.M{
// 			"jrs_kode": "$data.children.children.uk_kode",
// 			"jrs_nama": "$data.children.children.uk_nama",
// 		}},

// 		// Lanjut unwind ke prodi → metrics
// 		bson.M{"$unwind": "$data.children.children.children"},
// 		bson.M{"$unwind": "$data.children.children.children.metrics"},

// 		// Filter status
// 		bson.M{"$match": bson.M{
// 			"data.children.children.children.metrics.id_status_pegawai":   statusPegawai,
// 			"data.children.children.children.metrics.id_status_keaktifan": statusKeaktifan,
// 		}},

// 		// Group per jurusan
// 		bson.M{"$group": bson.M{
// 			"_id": bson.M{
// 				"jrs_kode": "$jrs_kode",
// 				"jrs_nama": "$jrs_nama",
// 			},
// 			"total": bson.M{"$sum": "$data.children.children.children.metrics.jumlah"},
// 		}},
// 		bson.M{"$sort": bson.M{"total": -1}},
// 	}

// 	cursor, err := repo.getCollectionByYear(tahun).Aggregate(ctx, pipeline)
// 	if err != nil {
// 		return nil, 0, err
// 	}
// 	defer cursor.Close(ctx)

// 	var results []bson.M
// 	if err = cursor.All(ctx, &results); err != nil {
// 		return nil, 0, err
// 	}

// 	var items []models.DrilldownItem
// 	var grandTotal int64

// 	for _, item := range results {
// 		id := item["_id"].(bson.M)
// 		value := utils.ConvertToInt64(item["total"])

// 		items = append(items, models.DrilldownItem{
// 			ID:    id["jrs_kode"].(string),
// 			Name:  id["jrs_nama"].(string),
// 			Value: value,
// 			Level: "jurusan",
// 		})
// 		grandTotal += value
// 	}

// 	return items, grandTotal, nil
// }

// // ------------------------------------------------------------
// // DRILLDOWN PRODI
// // Filter berdasarkan status + jurusan, baca metrics langsung
// // dari node prodi (level terdalam).
// // ------------------------------------------------------------
// func (repo *DashboardDosenMongoRepository) GetDrilldownDosenProdi(
// 	ctx context.Context, tahun, statusPegawai, statusKeaktifan int, kodeJurusan string,
// ) ([]models.DrilldownItem, int64, error) {

// 	pipeline := bson.A{
// 		bson.M{"$match": bson.M{"tahun": tahun}},

// 		bson.M{"$unwind": "$data"},
// 		bson.M{"$unwind": "$data.children"},
// 		bson.M{"$unwind": "$data.children.children"},

// 		// Filter hanya jurusan yang dipilih
// 		bson.M{"$match": bson.M{
// 			"data.children.children.uk_kode": kodeJurusan,
// 		}},

// 		// Unwind ke prodi → metrics
// 		bson.M{"$unwind": "$data.children.children.children"},
// 		bson.M{"$unwind": "$data.children.children.children.metrics"},

// 		// Filter status
// 		bson.M{"$match": bson.M{
// 			"data.children.children.children.metrics.id_status_pegawai":   statusPegawai,
// 			"data.children.children.children.metrics.id_status_keaktifan": statusKeaktifan,
// 		}},

// 		// Group per prodi — metrics sudah di level ini, tinggal baca
// 		bson.M{"$group": bson.M{
// 			"_id": bson.M{
// 				"prd_kode": "$data.children.children.children.uk_kode",
// 				"prd_nama": "$data.children.children.children.uk_nama",
// 			},
// 			"total": bson.M{"$sum": "$data.children.children.children.metrics.jumlah"},
// 		}},
// 		bson.M{"$sort": bson.M{"total": -1}},
// 	}

// 	cursor, err := repo.getCollectionByYear(tahun).Aggregate(ctx, pipeline)
// 	if err != nil {
// 		return nil, 0, err
// 	}
// 	defer cursor.Close(ctx)

// 	var results []bson.M
// 	if err = cursor.All(ctx, &results); err != nil {
// 		return nil, 0, err
// 	}

// 	var items []models.DrilldownItem
// 	var grandTotal int64

// 	for _, item := range results {
// 		id := item["_id"].(bson.M)
// 		value := utils.ConvertToInt64(item["total"])

// 		items = append(items, models.DrilldownItem{
// 			ID:    id["prd_kode"].(string),
// 			Name:  id["prd_nama"].(string),
// 			Value: value,
// 			Level: "prodi",
// 		})
// 		grandTotal += value
// 	}

// 	return items, grandTotal, nil
// }
