package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type AgendaMengajarService struct {
	AgendaMengajarRepository interfaces.AgendaMengajarRepository
}

func NewAgendaMengajarService(repo interfaces.AgendaMengajarRepository) *AgendaMengajarService {
	return &AgendaMengajarService{
		AgendaMengajarRepository: repo,
	}
}

func (service *AgendaMengajarService) GetAgendaMengajarFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.AgendaMengajar, int64, error) {
	return service.AgendaMengajarRepository.GetAgendaMengajarFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, page, limit)
}
