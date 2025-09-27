package interfaces

import (
	"context"
	"restapi-golang/models"
)

type RekapPMBRepository interface {
	GetRekapPMBFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, search string, tahun, page, limit int) ([]models.RekapPMB, int64, error)
}
