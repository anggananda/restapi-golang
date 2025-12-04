package mocks

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type BeasiswaMockRepository struct {
	mockData []models.Beasiswa
	err      error
}

func NewBeasiswaMockRepository(data []models.Beasiswa, err error) interfaces.BeasiswaRepository {
	return &BeasiswaMockRepository{
		mockData: data,
		err:      err,
	}
}

func (m *BeasiswaMockRepository) GetBeasiswaFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, semester, jenisBeasiswa, search string, tahun, page, limit int) ([]models.Beasiswa, int64, error) {
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
