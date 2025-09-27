package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type RekapPMBService struct {
	RekapPMBRepository interfaces.RekapPMBRepository
}

func NewRekapPMBService(repo interfaces.RekapPMBRepository) *RekapPMBService {
	return &RekapPMBService{
		RekapPMBRepository: repo,
	}
}

func (service *RekapPMBService) GetRekapPMBFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, search string, tahun, page, limit int) ([]models.RekapPMB, int64, error) {
	return service.RekapPMBRepository.GetRekapPMBFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, search, tahun, page, limit)
}
