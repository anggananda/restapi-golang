package mocks

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type RealisasiUnitMockRepository struct {
	mockData []models.RealisasiUnit
	err      error
}

func NewRealisasiUnitMockRepository(data []models.RealisasiUnit, err error) interfaces.RealisasiUnitRepository {
	return &RealisasiUnitMockRepository{
		mockData: data,
		err:      err,
	}
}

func (m *RealisasiUnitMockRepository) GetRealisasiUnitFiltered(ctx context.Context, search, tahun string, page, limit int) ([]models.RealisasiUnit, int64, error) {
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
