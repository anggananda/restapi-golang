package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type MhsWisudaService struct {
	MhsWisudaRepository interfaces.MhsWisudaRepository
}

func NewMhsWisudaService(repo interfaces.MhsWisudaRepository) *MhsWisudaService {
	return &MhsWisudaService{
		MhsWisudaRepository: repo,
	}
}

func (service *MhsWisudaService) GetMhsWisudaFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, search string, tahun, bulan, page, limit int) ([]models.MhsWisuda, int64, error) {
	return service.MhsWisudaRepository.GetMhsWisudaFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, search, tahun, bulan, page, limit)
}
