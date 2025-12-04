package mocks

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type BukuMockRepository struct {
	mockData []models.Buku
	err      error
}

func NewBukuMockRepository(data []models.Buku, err error) interfaces.BukuRepository {
	return &BukuMockRepository{
		mockData: data,
		err:      err,
	}
}

func (m *BukuMockRepository) GetBukuFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.Buku, int64, error) {
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
