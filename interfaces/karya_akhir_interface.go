package interfaces

import (
	"context"
	"restapi-golang/models"
)

type KaryaAkhirRepository interface {
	GetKaryaAkhirFiltered(ctx context.Context, kodefakultas, kodeJurusan, kodeProdi, search string, tahun, page, limit int)([]models.KaryaAkhir, int64, error)
}
