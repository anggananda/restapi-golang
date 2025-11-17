package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type KerjasamaService struct {
	KerjasamaRepository interfaces.KerjasamaRepository
}

func NewKerjasamaService(repo interfaces.KerjasamaRepository) *KerjasamaService {
	return &KerjasamaService{
		KerjasamaRepository: repo,
	}
}

func (service *KerjasamaService) GetKerjasamaFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, search string, page, limit int) ([]models.Kerjasama, int64, error) {
	return service.KerjasamaRepository.GetKerjasamaFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, search, page, limit)
}
