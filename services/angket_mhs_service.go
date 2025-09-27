package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type AngketMhsService struct {
	AngketMhsRepository interfaces.AngketMhsRepository
}

func NewAngketMhsService(repo interfaces.AngketMhsRepository) *AngketMhsService {
	return &AngketMhsService{
		AngketMhsRepository: repo,
	}
}

func (service *AngketMhsService) GetAngketMhsFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.AngketMhs, int64, error) {
	return service.AngketMhsRepository.GetAngketMhsFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, page, limit)
}
