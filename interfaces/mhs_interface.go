package interfaces

import (
	"context"
	"restapi-golang/models"
)

type MhsRepository interface {
	GetDetailMhs(ctx context.Context, nim string) (*models.Mahasiswa, error)
	GetMahasiswaHistoryFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, kewarganegaraan, search string, tahun, semester, angkatan, status, page, limit int) ([]models.MahasiswaHistoryResponse, int64, error)
}
