package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type KaryaAkhirService struct {
	KaryaAkhirRepository interfaces.KaryaAkhirRepository
}

func NewKaryaAkhirService(repo interfaces.KaryaAkhirRepository) *KaryaAkhirService {
	return &KaryaAkhirService{
		KaryaAkhirRepository: repo,
	}
}

func (service *KaryaAkhirService) GetKaryaAkhirFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, search string, tahun, page, limit int) ([]models.KaryaAkhir, int64, error) {
	return service.KaryaAkhirRepository.GetKaryaAkhirFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, search, tahun, page, limit)
}
