package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type RealisasiUnitService struct{
  RealisasiUnitRepository interfaces.RealisasiUnitRepository
}

func NewRealisasiUnitService(repo interfaces.RealisasiUnitRepository) *RealisasiUnitService{
  return &RealisasiUnitService{
    RealisasiUnitRepository: repo,
  }
}

func (service *RealisasiUnitService) GetRealisasiUnitFiltered(ctx context.Context, search string, page, limit int)([]models.RealisasiUnit, int64, error){
  return service.RealisasiUnitRepository.GetRealisasiUnitFiltered(ctx, search, page, limit)
}
