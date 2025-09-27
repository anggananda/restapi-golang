package interfaces

import (
	"context"
	"restapi-golang/models"
)

type PenelitianRepository interface {
	GetPenelitianFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.Penelitian, int64, error)
}
