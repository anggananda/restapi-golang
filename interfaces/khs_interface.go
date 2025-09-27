package interfaces

import (
	"context"
	"restapi-golang/models"
)

type KHSRepository interface{
  GetKHSFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int)([]models.Khs, int64, error)
}
