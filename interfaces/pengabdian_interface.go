package interfaces

import (
	"context"
	"restapi-golang/models"
)

type PengabdianRepository interface{
  GetPengabdianFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int)([]models.Pengabdian, int64, error)
}
