package interfaces

import (
	"context"
	"restapi-golang/models"
)

type StatusKeaktifanPegawaiRepository interface {
	GetStatusKeaktifanPegawai(ctx context.Context) ([]models.Status, error)
}
