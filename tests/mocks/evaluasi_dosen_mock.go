package mocks

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type EvaluasiDosenMockRepository struct {
	mockData []models.EvaluasiDosen
	err      error
}

func NewEvaluasiDosenMockRepository(data []models.EvaluasiDosen, err error) interfaces.EvaluasiDosenRepository {
	return &EvaluasiDosenMockRepository{
		mockData: data,
		err:      err,
	}
}

func (m *EvaluasiDosenMockRepository) GetEvaluasiDosenFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, namaDosen, search string, page, limit int) ([]models.EvaluasiDosen, int64, error) {
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
