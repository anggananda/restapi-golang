package interfaces

import (
	"context"
	"restapi-golang/models"
)

type StatusPegawaiRepository interface {
	GetStatusPegawai(ctx context.Context) ([]models.Status, error)
}
