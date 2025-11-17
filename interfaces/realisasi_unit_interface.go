package interfaces

import (
	"context"
	"restapi-golang/models"
)

type RealisasiUnitRepository interface {
	GetRealisasiUnitFiltered(ctx context.Context, search, tahun string, page, limit int) ([]models.RealisasiUnit, int64, error)
}
