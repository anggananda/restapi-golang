package mocks

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type HkiMockRepository struct {
	mockData []models.Hki
	err      error
}

func NewHkiMockRepository(data []models.Hki, err error) interfaces.HkiRepository {
	return &HkiMockRepository{
		mockData: data,
		err:      err,
	}
}

func (m *HkiMockRepository) GetHkiFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.Hki, int64, error) {
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
