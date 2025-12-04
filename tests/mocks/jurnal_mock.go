package mocks

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type JurnalMockRepository struct {
	mockData []models.Jurnal
	err      error
}

func NewJurnalMockRepository(data []models.Jurnal, err error) interfaces.JurnalRepository {
	return &JurnalMockRepository{
		mockData: data,
		err:      err,
	}
}

func (m *JurnalMockRepository) GetJurnalFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, indexer, akreditasi, search string, page, limit int) ([]models.Jurnal, int64, error) {
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
