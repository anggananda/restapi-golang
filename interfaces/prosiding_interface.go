package interfaces

import (
	"context"
	"restapi-golang/models"
)

type ProsidingRepository interface {
	GetProsidingFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, indexer, search string, page, limit int) ([]models.Prosiding, int64, error)
}
