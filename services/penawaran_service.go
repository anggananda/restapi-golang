package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type PenawaranService struct {
	PenawaranRepository interfaces.PenawaranRepository
}

func NewPenawaranService(repo interfaces.PenawaranRepository) *PenawaranService {
	return &PenawaranService{
		PenawaranRepository: repo,
	}
}

func (service *PenawaranService) GetPenawaranFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.Penawaran, int64, error) {
	return service.PenawaranRepository.GetPenawaranFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, page, limit)
}
