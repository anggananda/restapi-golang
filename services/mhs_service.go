package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type MhsService struct {
	MhsRepository interfaces.MhsRepository
}

func NewMhsService(repo interfaces.MhsRepository) *MhsService {
	return &MhsService{
		MhsRepository: repo,
	}
}

func (service *MhsService) GetDetailMhs(ctx context.Context, nim string) (*models.Mahasiswa, error) {
	return service.MhsRepository.GetDetailMhs(ctx, nim)
}

func (service *MhsService) GetMahasiswaHistoryFiltered(ctx context.Context,
	kodeFakultas, kodeJurusan, kodeProdi, kewarganegaraan, search string,
	tahun, semester int, angkatan string, status, page, limit int) ([]models.MahasiswaHistoryResponse, int64, error) {
	return service.MhsRepository.GetMahasiswaHistoryFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, kewarganegaraan, search, tahun, semester, angkatan, status, page, limit)
}
