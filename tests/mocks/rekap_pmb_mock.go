package mocks

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type RekapPMBMockRepository struct {
	mockData []models.RekapPMB
	err      error
}

func NewRekapPMBMockRepository(data []models.RekapPMB, err error) interfaces.RekapPMBRepository {
	return &RekapPMBMockRepository{
		mockData: data,
		err:      err,
	}
}

func (m *RekapPMBMockRepository) GetRekapPMBFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, search string, tahun, page, limit int) ([]models.RekapPMB, int64, error) {
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
