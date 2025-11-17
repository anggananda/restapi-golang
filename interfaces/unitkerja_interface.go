package interfaces

import (
	"context"
	"restapi-golang/models"
)

type UnitKerjaRepository interface {
	GetUnitKerja(ctx context.Context) (*models.UnitKerjaMapping, error)
}
