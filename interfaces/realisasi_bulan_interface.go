package interfaces

import (
	"context"
	"restapi-golang/models"
)

type RealisasiBulanRepository interface {
	GetRealisasiBulanFiltered(ctx context.Context, tahun, search string, page, limit int) ([]models.RealisasiBulan, int64, error)
}
