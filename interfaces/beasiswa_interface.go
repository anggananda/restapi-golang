package interfaces

import (
	"context"
	"restapi-golang/models"
)

type BeasiswaRepository interface{
  GetBeasiswaFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, semester, jenisBeasiswa, search string, tahun, page, limit int)([]models.Beasiswa, int64, error)
}
