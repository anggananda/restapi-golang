package interfaces

import (
	"context"
	"restapi-golang/models"
)

type AngketMhsRepository interface {
	GetAngketMhsFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.AngketMhs, int64, error)
}
