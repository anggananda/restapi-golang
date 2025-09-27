package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type BukuService struct {
	BukuRepository interfaces.BukuRepository
}

func NewBukuService(repo interfaces.BukuRepository) *BukuService {
	return &BukuService{
		BukuRepository: repo,
	}
}

func (service *BukuService) GetBukuFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.Buku, int64, error) {
	return service.BukuRepository.GetBukuFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, page, limit)
}
