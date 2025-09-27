package interfaces

import (
	"context"
	"restapi-golang/models"
)

type PerpemRepository interface {
	GetPerpemFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.Perpem, int64, error)
}
