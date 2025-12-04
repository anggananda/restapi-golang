package mocks

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type TracerMockRepository struct {
	mockData []models.Tracer
	err      error
}

func NewTracerMockRepository(data []models.Tracer, err error) interfaces.TracerRepository {
	return &TracerMockRepository{
		mockData: data,
		err:      err,
	}
}

func (m *TracerMockRepository) GetTracerFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, statusTracer, search string, tahun, bulan, page, limit int) ([]models.Tracer, int64, error) {
	if m.err != nil {
		return nil, 0, m.err
	}

	totalCount := int64(len(m.mockData))

	endIndex := page * limit
	if endIndex > int(totalCount) {
		endIndex = int(totalCount)
	}

	startIndex := (page - 1) * limit
	if startIndex < 0 {
		startIndex = 0
	}

	return m.mockData, totalCount, nil
}
