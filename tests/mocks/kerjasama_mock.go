package mocks

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type KerjasamaMockRepository struct {
	mockData []models.Kerjasama
	err      error
}

func NewKerjasamaMockRepository(data []models.Kerjasama, err error) interfaces.KerjasamaRepository {
	return &KerjasamaMockRepository{
		mockData: data,
		err:      err,
	}
}

func (m *KerjasamaMockRepository) GetKerjasamaFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, search string, page, limit int) ([]models.Kerjasama, int64, error) {
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
