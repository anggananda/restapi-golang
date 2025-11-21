package interfaces

import (
	"context"
	"restapi-golang/models"
)

type DashboardDosenRepository interface {
	GetDashboardDosenOverview(ctx context.Context, tahun int) ([]models.DashboardCardPegawai, error)
	GetDrilldownDosenFakultas(ctx context.Context, tahun, statusPegawai, statusKeaktifan int) ([]models.DrilldownItem, int64, error)
	GetDrilldownDosenJurusan(ctx context.Context, tahun, statusPegawai, statusKeaktifan int, kodeFakultas string) ([]models.DrilldownItem, int64, error)
	GetDrilldownDosenProdi(ctx context.Context, tahun, statusPegawai, statusKeaktifan int, kodeJurusan string) ([]models.DrilldownItem, int64, error)
}
