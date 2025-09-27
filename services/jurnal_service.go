package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type JurnalService struct {
	JurnalRepository interfaces.JurnalRepository
}

func NewJurnalService(repo interfaces.JurnalRepository) *JurnalService {
	return &JurnalService{
		JurnalRepository: repo,
	}
}

func (service *JurnalService) GetJurnalFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, indexer, akreditasi, search string, page, limit int) ([]models.Jurnal, int64, error) {
	return service.JurnalRepository.GetJurnalFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, indexer, akreditasi, search, page, limit)
}
