package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type PengabdianService struct {
	PengabdianRepository interfaces.PengabdianRepository
}

func NewPengabdianService(repo interfaces.PengabdianRepository) *PengabdianService {
	return &PengabdianService{
		PengabdianRepository: repo,
	}
}

func (service *PengabdianService) GetPengabdianFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.Pengabdian, int64, error) {
	return service.PengabdianRepository.GetPengabdianFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, page, limit)
}
