package interfaces

import (
	"context"
	"restapi-golang/models"
)

type DashboardMhsRepository interface {
	GetDashboardMhsOverview(ctx context.Context, tahun int, semester int) ([]models.DashboardCard, error)
	GetDrilldownMhsFakultas(ctx context.Context, tahun int, semester int, status string) ([]models.DrilldownItem, int64, error)
	GetDrilldownMhsJurusan(ctx context.Context, tahun int, semester int, status string, kodeFakultas string) ([]models.DrilldownItem, int64, error)
	GetDrilldownMhsProdi(ctx context.Context, tahun int, semester int, status string, kodeJurusan string) ([]models.DrilldownItem, int64, error)
	HasJurusan(ctx context.Context, tahun int, semester int, status string, kodeFakultas string) (bool, error)
	GetProdiByFakultas(ctx context.Context, tahun int, semester int, status string, kodeFakultas string) ([]models.DrilldownItem, int64, error)
}
