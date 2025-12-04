package mocks

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type AgendaMengajarMockRepository struct {
	mockData []models.AgendaMengajar
	err      error
}

func NewAgendaMengajarMockRepository(data []models.AgendaMengajar, err error) interfaces.AgendaMengajarRepository {
	return &AgendaMengajarMockRepository{
		mockData: data,
		err:      err,
	}
}

func (m *AgendaMengajarMockRepository) GetAgendaMengajarFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.AgendaMengajar, int64, error) {
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
