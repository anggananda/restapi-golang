package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type EvaluasiDosenService struct {
	EvaluasiDosenRepository interfaces.EvaluasiDosenRepository
}

func NewEvaluasiDosenService(repo interfaces.EvaluasiDosenRepository) *EvaluasiDosenService {
	return &EvaluasiDosenService{
		EvaluasiDosenRepository: repo,
	}
}

func (service *EvaluasiDosenService) GetEvaluasiDosenFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, namaDosen, search string, page, limit int) ([]models.EvaluasiDosen, int64, error) {
	return service.EvaluasiDosenRepository.GetEvaluasiDosenFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, namaDosen, search, page, limit)
}
