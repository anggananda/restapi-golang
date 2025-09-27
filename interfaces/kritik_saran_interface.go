package interfaces

import (
	"context"
	"restapi-golang/models"
)

type KritikSaranRepository interface {
	GetKritikSaranFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.KritikSaran, int64, error)
}
