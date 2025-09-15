package interfaces

import (
	"context"
	"restapi-golang/models"
)

type PerpemRepository interface {
	GetPerpemFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester string) ([]models.Perpem, error)
}
