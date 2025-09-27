package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type ProsidingService struct {
	ProsidingRepository interfaces.ProsidingRepository
}

func NewProsidingService(repo interfaces.ProsidingRepository) *ProsidingService {
	return &ProsidingService{
		ProsidingRepository: repo,
	}
}

func (service *ProsidingService) GetProsidingFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, indexer, search string, page, limit int) ([]models.Prosiding, int64, error) {
	return service.ProsidingRepository.GetProsidingFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, indexer, search, page, limit)
}
