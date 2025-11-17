package interfaces

import (
	"context"
	"restapi-golang/models"
)

type KerjasamaRepository interface {
	GetKerjasamaFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, search string, page, limit int) ([]models.Kerjasama, int64, error)
}
