package interfaces

import (
	"context"
	"restapi-golang/models"
)

type DashboardMhsRepository interface {
	GetDashboardOverview(ctx context.Context, tahun int, semesterType string) ([]models.DashboardCard, error)
	GetDrilldownFakultas(ctx context.Context, tahun int, semesterType string, status string) ([]models.DrilldownItem, int64, error)
	GetDrilldownJurusan(ctx context.Context, tahun int, semesterType string, status string, fakultasKode string) ([]models.DrilldownItem, int64, error)
	GetDrilldownProdi(ctx context.Context, tahun int, semesterType string, status string, jurusanKode string) ([]models.DrilldownItem, int64, error)
}
