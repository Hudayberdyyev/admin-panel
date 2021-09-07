package service

import (
	"github.com/Hudayberdyyev/admin-panel-backend/models"
	"github.com/Hudayberdyyev/admin-panel-backend/pkg/repository"
)

type ExtraService struct {
	repos repository.ExtraMessages
}

func NewExtraService(repos repository.ExtraMessages) *ExtraService {
	return &ExtraService{repos: repos}
}

func (s *ExtraService) Create(input models.ExtraMessages) (int, error) {
	return s.repos.Create(input)
}

func (s *ExtraService) Delete(id int) error {
	return s.repos.Delete(id)
}

func (s *ExtraService) GetAll() ([]models.ExtraMessages, error) {
	return s.repos.GetAll()
}
