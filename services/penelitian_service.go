package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type PenelitianService struct {
	PenelitianRepository interfaces.PenelitianRepository
}

func NewPenelitianService(repo interfaces.PenelitianRepository) *PenelitianService {
	return &PenelitianService{
		PenelitianRepository: repo,
	}
}

func (service *PenelitianService) GetPenelitianFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.Penelitian, int64, error) {
	return service.PenelitianRepository.GetPenelitianFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, page, limit)
}
