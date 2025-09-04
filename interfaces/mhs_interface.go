package interfaces

import (
	"context"
	"restapi-golang/models"
)

type MhsRepository interface {
	GetDetailMhs(ctx context.Context, nim string) (*models.Mahasiswa, error)
	GetMahasiswaHistoryByStatus(ctx context.Context, status string, page, limit int, tahun int, semesterType string) ([]models.MahasiswaHistoryResponse, int64, error)
	GetMahasiswaHistoryFiltered(ctx context.Context, filter models.MahasiswaHistoryRequest) ([]models.MahasiswaHistoryResponse, int64, error)
}
