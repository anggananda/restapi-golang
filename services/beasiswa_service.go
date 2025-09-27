package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type BeasiswaService struct {
	BeasiswaRepository interfaces.BeasiswaRepository
}

func NewBeasiswaService(repo interfaces.BeasiswaRepository) *BeasiswaService {
	return &BeasiswaService{
		BeasiswaRepository: repo,
	}
}

func (service *BeasiswaService) GetBeasiswaFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, semester, jenisBeasiswa, search string, tahun, page, limit int) ([]models.Beasiswa, int64, error) {
	return service.BeasiswaRepository.GetBeasiswaFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, semester, jenisBeasiswa, search, tahun, page, limit)
}
