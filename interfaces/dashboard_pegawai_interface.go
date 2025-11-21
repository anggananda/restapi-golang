package interfaces

import (
	"context"
	"restapi-golang/models"
)

type DashboardPegawaiRepository interface {
	GetDashboardPegawaiOverview(ctx context.Context, tahun int) ([]models.DashboardCardPegawai, error)
	GetDrilldownPegawaiFakultas(ctx context.Context, tahun, statusPegawai, statusKeaktifan int) ([]models.DrilldownItem, int64, error)
	GetDrilldownPegawaiJurusan(ctx context.Context, tahun, statusPegawai, statusKeaktifan int, kodeFakultas string) ([]models.DrilldownItem, int64, error)
	GetDrilldownPegawaiProdi(ctx context.Context, tahun, statusPegawai, statusKeaktifan int, kodeJurusan string) ([]models.DrilldownItem, int64, error)
}
