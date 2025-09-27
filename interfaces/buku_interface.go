package interfaces

import (
	"context"
	"restapi-golang/models"
)

type BukuRepository interface {
	GetBukuFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.Buku, int64, error)
}
