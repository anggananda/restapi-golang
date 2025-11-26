package interfaces

import (
	"context"
	"restapi-golang/models"
)

type DosenRepository interface {
	GetDetailDosen(ctx context.Context, niu string) (*models.Dosen, error)
	GetDosenHistoryFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, search string, tahun, statusPegawai, statusKeaktifan, page, limit int) ([]models.DosenHistoryResponse, int64, error)
}
