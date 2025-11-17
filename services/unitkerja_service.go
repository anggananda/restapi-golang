package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type UnitKerjaService struct {
	UnitKerjaRepository interfaces.UnitKerjaRepository
}

func NewUnitKerjaService(repo interfaces.UnitKerjaRepository) *UnitKerjaService {
	return &UnitKerjaService{
		UnitKerjaRepository: repo,
	}
}

func (service *UnitKerjaService) GetUnitKerja(ctx context.Context) (*models.UnitKerjaMapping, error) {
	return service.UnitKerjaRepository.GetUnitKerja(ctx)
}
