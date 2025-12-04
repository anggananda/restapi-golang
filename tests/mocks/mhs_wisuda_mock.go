package mocks

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type MhsWisudaMockRepository struct {
	mockData []models.MhsWisuda
	err      error
}

func NewMhsWisudaMockRepository(data []models.MhsWisuda, err error) interfaces.MhsWisudaRepository {
	return &MhsWisudaMockRepository{
		mockData: data,
		err:      err,
	}
}

func (m *MhsWisudaMockRepository) GetMhsWisudaFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, search string, tahun, bulan, page, limit int) ([]models.MhsWisuda, int64, error) {
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
