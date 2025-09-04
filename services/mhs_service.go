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

func (service *MhsService) GetMahasiswaHistoryByStatus(ctx context.Context, status string, page, limit int, tahun int, semesterType string) ([]models.MahasiswaHistoryResponse, int64, error) {
	return service.MhsRepository.GetMahasiswaHistoryByStatus(ctx, status, page, limit, tahun, semesterType)
}

func (service *MhsService) GetMahasiswaHistoryFiltered(ctx context.Context, filter models.MahasiswaHistoryRequest) ([]models.MahasiswaHistoryResponse, int64, error) {
	return service.MhsRepository.GetMahasiswaHistoryFiltered(ctx, filter)
}
