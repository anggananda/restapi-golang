package interfaces

import (
	"context"
	"restapi-golang/models"
)

type PegawaiRepository interface {
	GetDetailPegawai(ctx context.Context, niu string) (*models.PegawaiHistory, error)
	GetPegawaiHistoryFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, search string, tahun, statusPegawai, statusKeaktifan, page, limit int) ([]models.PegawaiHistoryResponse, int64, error)
}
