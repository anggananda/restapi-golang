package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type HkiService struct {
	HkiRepository interfaces.HkiRepository
}

func NewHkiService(repo interfaces.HkiRepository) *HkiService {
	return &HkiService{
		HkiRepository: repo,
	}
}

func (service *HkiService) GetHkiFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search string, page, limit int) ([]models.Hki, int64, error) {
	return service.HkiRepository.GetHkiFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, page, limit)
}
