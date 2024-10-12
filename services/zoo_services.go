package services

import (
    "zoo-backend/models"
    "zoo-backend/repositories"
)

type ZooService struct {
    Repo *repositories.ZooRepository
}

func (s *ZooService) CreateZoo(zoo models.Zoo) (int64, error) {
    return s.Repo.Create(zoo)
}

func (s *ZooService) GetAllZoos() ([]models.Zoo, error) {
    return s.Repo.GetAll()
}

func (s *ZooService) GetZooByID(id int) (models.Zoo, error) {
    return s.Repo.GetByID(id)
}

func (s *ZooService) UpdateZoo(zoo models.Zoo) error {
    return s.Repo.Update(zoo)
}

func (s *ZooService) DeleteZoo(id int) error {
    return s.Repo.Delete(id)
}
