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

func (service *DashboardMhsService) GetDashboardMhsOverview(ctx context.Context, tahun int, semester int) ([]models.DashboardCard, error) {
	return service.DashboardMhsRepository.GetDashboardMhsOverview(ctx, tahun, semester)
}

func (service *DashboardMhsService) GetDrilldownMhsFakultas(ctx context.Context, tahun int, semester int, status string) ([]models.DrilldownItem, int64, error) {
	return service.DashboardMhsRepository.GetDrilldownMhsFakultas(ctx, tahun, semester, status)
}

func (service *DashboardMhsService) GetDrilldownMhsJurusan(ctx context.Context, tahun int, semester int, status string, kodeFakultas string) ([]models.DrilldownItem, int64, error) {
	return service.DashboardMhsRepository.GetDrilldownMhsJurusan(ctx, tahun, semester, status, kodeFakultas)
}

func (service *DashboardMhsService) GetDrilldownMhsProdi(ctx context.Context, tahun int, semester int, status string, kodeJurusan string) ([]models.DrilldownItem, int64, error) {
	return service.DashboardMhsRepository.GetDrilldownMhsProdi(ctx, tahun, semester, status, kodeJurusan)
}
