package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type DosenService struct {
	DosenRepository interfaces.DosenRepository
}

func NewDosenService(repo interfaces.DosenRepository) *DosenService {
	return &DosenService{
		DosenRepository: repo,
	}
}

func (service *DosenService) GetDetailDosen(ctx context.Context, niu string) (*models.Dosen, error) {
	return service.DosenRepository.GetDetailDosen(ctx, niu)
}

func (service *DosenService) GetDosenHistoryFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, search string, tahun, statusPegawai, statusKeaktifan, page, limit int) ([]models.DosenHistoryResponse, int64, error) {
	return service.DosenRepository.GetDosenHistoryFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, search, tahun, statusPegawai, statusKeaktifan, page, limit)
}
