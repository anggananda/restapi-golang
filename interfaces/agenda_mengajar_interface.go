package interfaces

import (
	"context"
	"restapi-golang/models"
)

type AgendaMengajarRepository interface{
  GetAgendaMengajarFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int)([]models.AgendaMengajar, int64, error)
}
