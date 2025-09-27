package interfaces

import (
	"context"
	"restapi-golang/models"
)

type MhsWisudaRepository interface {
	GetMhsWisudaFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, search string, tahun, bulan, page, limit int) ([]models.MhsWisuda, int64, error)
}
