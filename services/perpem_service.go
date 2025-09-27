package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type PerpemService struct {
	PerpemRepository interfaces.PerpemRepository
}

func NewPerpemService(repo interfaces.PerpemRepository) *PerpemService {
	return &PerpemService{
		PerpemRepository: repo,
	}
}

func (service *PerpemService) GetPerpemFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.Perpem, int64, error) {
	return service.PerpemRepository.GetPerpemFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, page, limit)
}
