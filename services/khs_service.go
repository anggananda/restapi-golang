package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type KHSService struct {
	KHSRepository interfaces.KHSRepository
}

func NewKHSService(repo interfaces.KHSRepository) *KHSService {
	return &KHSService{
		KHSRepository: repo,
	}
}

func (service *KHSService) GetKHSFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.Khs, int64, error) {
	return service.KHSRepository.GetKHSFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, page, limit)
}
