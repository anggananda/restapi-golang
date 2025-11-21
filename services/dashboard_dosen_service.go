package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type DashboardDosenService struct {
	DashboardDosenRepository interfaces.DashboardDosenRepository
}

func NewDashboardDosenService(repo interfaces.DashboardDosenRepository) *DashboardDosenService {
	return &DashboardDosenService{
		DashboardDosenRepository: repo,
	}
}

func (service *DashboardDosenService) GetDashboardDosenOverview(ctx context.Context, tahun int) ([]models.DashboardCardPegawai, error) {
	return service.DashboardDosenRepository.GetDashboardDosenOverview(ctx, tahun)
}

func (service *DashboardDosenService) GetDrilldownDosenFakultas(ctx context.Context, tahun, statusPegawai, statusKeaktifan int) ([]models.DrilldownItem, int64, error) {
	return service.DashboardDosenRepository.GetDrilldownDosenFakultas(ctx, tahun, statusPegawai, statusKeaktifan)
}

func (service *DashboardDosenService) GetDrilldownDosenJurusan(ctx context.Context, tahun, statusPegawai, statusKeaktifan int, kodeFakultas string) ([]models.DrilldownItem, int64, error) {
	return service.DashboardDosenRepository.GetDrilldownDosenJurusan(ctx, tahun, statusPegawai, statusKeaktifan, kodeFakultas)
}

func (service *DashboardDosenService) GetDrilldownDosenProdi(ctx context.Context, tahun, statusPegawai, statusKeaktifan int, kodeJurusan string) ([]models.DrilldownItem, int64, error) {
	return service.DashboardDosenRepository.GetDrilldownDosenProdi(ctx, tahun, statusPegawai, statusKeaktifan, kodeJurusan)
}
