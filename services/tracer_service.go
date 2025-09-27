package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type TracerService struct {
	TracerRepository interfaces.TracerRepository
}

func NewTracerService(repo interfaces.TracerRepository) *TracerService {
	return &TracerService{
		TracerRepository: repo,
	}
}

func (service *TracerService) GetTracerFiltered(ctx context.Context, kodeFakultas, kodeJurusan, kodeProdi, statusTracer, search string, tahun, bulan, page, limit int) ([]models.Tracer, int64, error){
  return service.TracerRepository.GetTracerFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, statusTracer, search, tahun, bulan, page, limit)
}
