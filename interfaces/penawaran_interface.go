package interfaces

import (
	"context"
	"restapi-golang/models"
)

type PenawaranRepository interface {
	GetPenawaranFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.Penawaran, int64, error)
}
