package interfaces

import (
	"context"
	"restapi-golang/models"
)

type StatusMahasiswaRepository interface {
	GetStatusMahasiswa(ctx context.Context) ([]models.Status, error)
}
