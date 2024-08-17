package service

import (
	"github.com/scostadavid/websteady/internal/app/monitorable/dto"
	"github.com/scostadavid/websteady/internal/app/monitorable/entity"
	"github.com/scostadavid/websteady/internal/app/monitorable/repository"
)

// entidade é usada pelo repositório que é usada pelo serviço que é usada pela main, outros serviços só se conversam via broker
type MonitorableService struct {
	repository repository.ReaderWriter
}

func NewMonitorableService(repository repository.ReaderWriter) *MonitorableService {
	return &MonitorableService{
		repository: repository,
	}
}

func (s *MonitorableService) GetAllMonitorables() ([]*entity.Monitorable, error) {
	monitorables, _ := s.repository.GetAll()
	return monitorables, nil
}

func (s *MonitorableService) AddMonitorable(dto *dto.AddMonitorable) (*entity.Monitorable, error) {
	monitorable, err := s.repository.Create(dto)

	if err != nil {
		return nil, err
	}

	return monitorable, nil
}
