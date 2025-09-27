package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type RealisasiBulanService struct {
	RealisasiBulanRepository interfaces.RealisasiBulanRepository
}

func NewRealisasiBulanService(repo interfaces.RealisasiBulanRepository) *RealisasiBulanService {
	return &RealisasiBulanService{
		RealisasiBulanRepository: repo,
	}
}

func (service *RealisasiBulanService) GetRealisasiBulanFiltered(ctx context.Context, tahun, search string, page, limit int) ([]models.RealisasiBulan, int64, error) {
	return service.RealisasiBulanRepository.GetRealisasiBulanFiltered(ctx, tahun, search, page, limit)
}
