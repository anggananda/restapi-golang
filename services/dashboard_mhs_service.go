package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type DashboardMhsService struct {
	DashboardMhsRepository interfaces.DashboardMhsRepository
}

func NewDashboardMhsService(repo interfaces.DashboardMhsRepository) *DashboardMhsService {
	return &DashboardMhsService{
		DashboardMhsRepository: repo,
	}
}

func (service *DashboardMhsService) GetDashboardOverview(ctx context.Context, tahun int, semester string) ([]models.DashboardCard, error) {
	return service.DashboardMhsRepository.GetDashboardOverview(ctx, tahun, semester)
}

func (service *DashboardMhsService) GetDrilldownFakultas(ctx context.Context, tahun int, semester string, status string) ([]models.DrilldownItem, int64, error) {
	return service.DashboardMhsRepository.GetDrilldownFakultas(ctx, tahun, semester, status)
}

func (service *DashboardMhsService) GetDrilldownJurusan(ctx context.Context, tahun int, semester string, status string, fakultasKode string) ([]models.DrilldownItem, int64, error) {
	return service.DashboardMhsRepository.GetDrilldownJurusan(ctx, tahun, semester, status, fakultasKode)
}

func (service *DashboardMhsService) GetDrilldownProdi(ctx context.Context, tahun int, semester string, status string, jurusanKode string) ([]models.DrilldownItem, int64, error) {
	return service.DashboardMhsRepository.GetDrilldownProdi(ctx, tahun, semester, status, jurusanKode)
}
