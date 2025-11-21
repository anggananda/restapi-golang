package interfaces

import (
	"context"
	"restapi-golang/models"
)

type EvaluasiDosenRepository interface {
	GetEvaluasiDosenFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, namaDosen, search string, page, limit int) ([]models.EvaluasiDosen, int64, error)
}
