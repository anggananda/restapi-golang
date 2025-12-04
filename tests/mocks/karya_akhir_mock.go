package mocks

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type KaryaAkhirMockRepository struct {
	mockData []models.KaryaAkhir
	err      error
}

func NewKaryaAkhirMockRepository(data []models.KaryaAkhir, err error) interfaces.KaryaAkhirRepository {
	return &KaryaAkhirMockRepository{
		mockData: data,
		err:      err,
	}
}

func (m *KaryaAkhirMockRepository) GetKaryaAkhirFiltered(ctx context.Context, kodefakultas, kodeJurusan, kodeProdi, search string, tahun, page, limit int) ([]models.KaryaAkhir, int64, error) {
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
