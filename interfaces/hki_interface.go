package interfaces

import (
	"context"
	"restapi-golang/models"
)

type HkiRepository interface {
	GetHkiFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.Hki, int64, error)
}
