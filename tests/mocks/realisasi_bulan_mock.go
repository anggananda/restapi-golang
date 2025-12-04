package mocks

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type RealisasiBulanMockRepository struct {
	mockData []models.RealisasiBulan
	err      error
}

func NewRealisasiBulanMockRepository(data []models.RealisasiBulan, err error) interfaces.RealisasiBulanRepository {
	return &RealisasiBulanMockRepository{
		mockData: data,
		err:      err,
	}
}

func (m *RealisasiBulanMockRepository) GetRealisasiBulanFiltered(ctx context.Context, tahun, search string, page, limit int) ([]models.RealisasiBulan, int64, error) {
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
