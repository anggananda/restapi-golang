package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type DashboardPegawaiService struct {
	DashboardPegawaiRepository interfaces.DashboardPegawaiRepository
}

func NewDashboardPegawaiService(repo interfaces.DashboardPegawaiRepository) *DashboardPegawaiService {
	return &DashboardPegawaiService{
		DashboardPegawaiRepository: repo,
	}
}

func (service *DashboardPegawaiService) GetDashboardPegawaiOverview(ctx context.Context, tahun int) ([]models.DashboardCardPegawai, error) {
	return service.DashboardPegawaiRepository.GetDashboardPegawaiOverview(ctx, tahun)
}

func (service *DashboardPegawaiService) GetDrilldownPegawaiFakultas(ctx context.Context, tahun, statusPegawai, statusKeaktifan int) ([]models.DrilldownItem, int64, error) {
	return service.DashboardPegawaiRepository.GetDrilldownPegawaiFakultas(ctx, tahun, statusPegawai, statusKeaktifan)
}

func (service *DashboardPegawaiService) GetDrilldownPegawaiJurusan(ctx context.Context, tahun, statusPegawai, statusKeaktifan int, kodeFakultas string) ([]models.DrilldownItem, int64, error) {
	return service.DashboardPegawaiRepository.GetDrilldownPegawaiJurusan(ctx, tahun, statusPegawai, statusKeaktifan, kodeFakultas)
}

func (service *DashboardPegawaiService) GetDrilldownPegawaiProdi(ctx context.Context, tahun, statusPegawai, statusKeaktifan int, kodeJurusan string) ([]models.DrilldownItem, int64, error) {
	return service.DashboardPegawaiRepository.GetDrilldownPegawaiProdi(ctx, tahun, statusPegawai, statusKeaktifan, kodeJurusan)
}
