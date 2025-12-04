package mocks

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type ProsidingMockRepository struct {
	mockData []models.Prosiding
	err      error
}

func NewProsidingMockRepository(data []models.Prosiding, err error) interfaces.ProsidingRepository {
	return &ProsidingMockRepository{
		mockData: data,
		err:      err,
	}
}

func (m *ProsidingMockRepository) GetProsidingFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, indexer, search string, page, limit int) ([]models.Prosiding, int64, error) {
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
