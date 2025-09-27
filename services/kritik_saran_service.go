package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type KritikSaranService struct {
	KritikSaranRepository interfaces.KritikSaranRepository
}

func NewKritikSaranService(repo interfaces.KritikSaranRepository) *KritikSaranService {
	return &KritikSaranService{
		KritikSaranRepository: repo,
	}
}

func (service *KritikSaranService) GetKritikSaranFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.KritikSaran, int64, error) {
	return service.KritikSaranRepository.GetKritikSaranFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, page, limit)
}
