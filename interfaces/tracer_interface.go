package interfaces

import (
	"context"
	"restapi-golang/models"
)

type TracerRepository interface {
	GetTracerFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, statusTracer, search string, tahun, bulan, page, limit int) ([]models.Tracer, int64, error)
}
