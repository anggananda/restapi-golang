package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type PegawaiService struct {
	PegawaiRepository interfaces.PegawaiRepository
}

func NewPegawaiService(repo interfaces.PegawaiRepository) *PegawaiService {
	return &PegawaiService{
		PegawaiRepository: repo,
	}
}

func (service *PegawaiService) GetDetailPegawai(ctx context.Context, niu string) (*models.PegawaiHistory, error) {
return service.PegawaiRepository.GetDetailPegawai(ctx, niu)
}

func (service *PegawaiService) GetPegawaiHistoryFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, search string, tahun, statusPegawai, statusKeaktifan, page, limit int) ([]models.PegawaiHistoryResponse, int64, error) {
	return service.PegawaiRepository.GetPegawaiHistoryFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, search, tahun, statusPegawai, statusKeaktifan, page, limit)
}
