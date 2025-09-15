package repositories

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
	"restapi-golang/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MhsRepository struct {
	Collection *mongo.Collection
}

func NewMhsRepository(db *mongo.Database) interfaces.MhsRepository {
	return &MhsRepository{
		Collection: db.Collection("mahasiswa_v3"),
	}
}

func (repo *MhsRepository) GetDetailMhs(ctx context.Context, nim string) (*models.Mahasiswa, error) {
	var mh models.Mahasiswa

	if err := repo.Collection.FindOne(ctx, bson.M{"nim": nim}).Decode(&mh); err != nil {
		return nil, err
	}

	return &mh, nil
}

func (repo *MhsRepository) GetMahasiswaHistoryByStatus(ctx context.Context, status string, page, limit int, tahun int, semesterType string) ([]models.MahasiswaHistoryResponse, int64, error) {
	skip := (page - 1) * limit

	pipeline := bson.A{
		// Match dokumen yang memiliki history sesuai kriteria
		bson.M{"$match": bson.M{
			"history.tahun":         tahun,
			"history.semester_type": semesterType,
			"history.status":        status,
		}},

		// Project dengan filter history
		bson.M{"$project": bson.M{
			"nim":             1,
			"nama":            1,
			"tahun_masuk":     1,
			"kewarganegaraan": 1,
			"fakultas":        "$unit.fakultas",
			"jurusan":         "$unit.jurusan",
			"prodi":           "$unit.prodi",

			// Filter history yang sesuai kriteria
			"filtered_history": bson.M{
				"$filter": bson.M{
					"input": "$history",
					"as":    "h",
					"cond": bson.M{
						"$and": bson.A{
							bson.M{"$eq": bson.A{"$$h.tahun", tahun}},
							bson.M{"$eq": bson.A{"$$h.semester_type", semesterType}},
							bson.M{"$eq": bson.A{"$$h.status", status}},
						},
					},
				},
			},
		}},

		// Unwind hanya history yang sudah difilter
		bson.M{"$unwind": "$filtered_history"},

		// Facet untuk data dan total
		bson.M{"$facet": bson.M{
			"data": bson.A{
				bson.M{"$project": bson.M{
					"nim":              "$nim",
					"nama":             "$nama",
					"tahun_masuk":      "$tahun_masuk",
					"kewarganegaraan":  "$kewarganegaraan",
					"fakultas":         "$fakultas",
					"jurusan":          "$jurusan",
					"prodi":            "$prodi",
					"tahun":            "$filtered_history.tahun",
					"semester":         "$filtered_history.semester",
					"status":           "$filtered_history.status",
					"status_singkatan": "$filtered_history.status_singkatan",
					"nama_pa":          "$filtered_history.nama_pa",
				}},
				bson.M{"$sort": bson.M{"nim": 1, "tahun": 1, "semester": 1}},
				bson.M{"$skip": skip},
				bson.M{"$limit": limit},
			},
			"total": bson.A{
				bson.M{"$count": "total"},
			},
		}},
	}

	cursor, err := repo.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	// Parse hasil
	var facetResult []struct {
		Data  []models.MahasiswaHistoryResponse `bson:"data"`
		Total []struct {
			Total int64 `bson:"total"`
		} `bson:"total"`
	}

	if err = cursor.All(ctx, &facetResult); err != nil {
		return nil, 0, err
	}

	// Handle hasil
	var results []models.MahasiswaHistoryResponse
	var total int64

	if len(facetResult) > 0 {
		results = facetResult[0].Data
		if len(facetResult[0].Total) > 0 {
			total = facetResult[0].Total[0].Total
		}
	}

	return results, total, nil
}

func (repo *MhsRepository) GetMahasiswaHistoryFiltered(ctx context.Context, filter models.MahasiswaHistoryRequest) ([]models.MahasiswaHistoryResponse, int64, error) {
	matchConditions := bson.M{
		"history.tahun":         filter.Tahun,
		"history.semester_type": filter.Semester,
	}

	if filter.Nama != "" {
		matchConditions["nama"] = filter.Nama
	}
	if filter.Fakultas != "" {
		matchConditions["unit.fakultas"] = filter.Fakultas
	}
	if filter.Jurusan != "" {
		matchConditions["unit.jurusan"] = filter.Jurusan
	}
	if filter.Prodi != "" {
		matchConditions["unit.prodi"] = filter.Prodi
	}
	if filter.Status != "" {
		matchConditions["history.status"] = filter.Status
	}
	if filter.Kewarganegaraan != "" {
		matchConditions["kewarganegaraan"] = filter.Kewarganegaraan
	}
	if filter.NIM != "" {
		matchConditions["nim"] = filter.NIM
	}

	skip := (filter.Page - 1) * filter.Limit

	pipeline := bson.A{
		bson.M{"$unwind": "$history"},
		bson.M{"$match": matchConditions},
		bson.M{"$project": bson.M{
			"nim":              1,
			"nama":             1,
			"tahun_masuk":      1,
			"kewarganegaraan":  1,
			"fakultas":         "$unit.fakultas",
			"jurusan":          "$unit.jurusan",
			"prodi":            "$unit.prodi",
			"tahun":            "$history.tahun",
			"semester":         "$history.semester",
			"status":           "$history.status",
			"status_singkatan": "$history.status_singkatan",
			"nama_pa":          "$history.nama_pa",
		}},
		bson.M{"$sort": bson.M{"nim": 1, "tahun": 1, "semester": 1}},
		bson.M{"$skip": skip},
		bson.M{"$limit": filter.Limit},
	}

	cursor, err := repo.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var results []models.MahasiswaHistoryResponse
	if err = cursor.All(ctx, &results); err != nil {
		return nil, 0, err
	}

	countPipeline := bson.A{
		bson.M{"$unwind": "$history"},
		bson.M{"$match": matchConditions},
		bson.M{"$count": "total"},
	}

	countCursor, err := repo.Collection.Aggregate(ctx, countPipeline)
	if err != nil {
		return nil, 0, err
	}
	defer countCursor.Close(ctx)

	var countResult []bson.M
	if err = countCursor.All(ctx, &countResult); err != nil {
		return nil, 0, err
	}

	var total int64
	if len(countResult) > 0 {
		total = utils.ConvertToInt64(countResult[0]["total"])
	}

	return results, total, nil
}
