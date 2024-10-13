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

func (s *ZooService) UpsertZoo(zoo models.Zoo) (bool, error) {
    return s.Repo.Upsert(zoo)
}



func (s *ZooService) DeleteZoo(id int) error {
    existingZoo, err := s.Repo.GetByID(id) 
    if err != nil {
        return err 
    }

    return s.Repo.Delete(existingZoo.ID) 
}
