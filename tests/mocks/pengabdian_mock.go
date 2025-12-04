package mocks

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type PengabdianMockRepository struct {
	mockData []models.Pengabdian
	err      error
}

func NewPengabdianMockRepository(data []models.Pengabdian, err error) interfaces.PengabdianRepository {
	return &PengabdianMockRepository{
		mockData: data,
		err:      err,
	}
}

func (m *PengabdianMockRepository) GetPengabdianFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.Pengabdian, int64, error) {
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
