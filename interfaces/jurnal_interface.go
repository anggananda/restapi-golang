package interfaces

import (
	"context"
	"restapi-golang/models"
)

type JurnalRepository interface{
  GetJurnalFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, indexer, akreditasi, search string, page, limit int)([]models.Jurnal, int64, error)
}
