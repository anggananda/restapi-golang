package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type StatusService struct {
	StatusMahasiswaRepository        interfaces.StatusMahasiswaRepository
	StatusPegawaiRepository          interfaces.StatusPegawaiRepository
	StatusKeaktifanPegawaiRepository interfaces.StatusKeaktifanPegawaiRepository
}

func NewStatusService(repoStatusMhs interfaces.StatusMahasiswaRepository, repoStatusPegawai interfaces.StatusPegawaiRepository, repoStatusKeaktifanPegawai interfaces.StatusKeaktifanPegawaiRepository) *StatusService {
	return &StatusService{
		StatusMahasiswaRepository:        repoStatusMhs,
		StatusPegawaiRepository:          repoStatusPegawai,
		StatusKeaktifanPegawaiRepository: repoStatusKeaktifanPegawai,
	}
}

func (repo *StatusService) GetStatusMahasiswa(ctx context.Context) ([]models.Status, error) {
	return repo.StatusMahasiswaRepository.GetStatusMahasiswa(ctx)
}
func (repo *StatusService) GetStatusPegawai(ctx context.Context) ([]models.Status, error) {
	return repo.StatusPegawaiRepository.GetStatusPegawai(ctx)
}
func (repo *StatusService) GetStatusKeaktifanPegawai(ctx context.Context) ([]models.Status, error) {
	return repo.StatusKeaktifanPegawaiRepository.GetStatusKeaktifanPegawai(ctx)
}
