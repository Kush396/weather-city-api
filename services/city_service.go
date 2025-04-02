package services

import (
	"weather-city-api/models"
	"weather-city-api/repository"
)

type CityService interface {
	GetAllCities() []models.City
	GetCityByID(id int) (models.City, error)
	CreateCity(city models.City) (models.City, error)
	UpdateCity(id int, city models.City) (models.City, error)
	DeleteCity(id int) error
}

type cityService struct {
	repo repository.CityRepository
}

func NewCityService(repo repository.CityRepository) CityService {
	return &cityService{
		repo: repo,
	}
}

func (s *cityService) GetAllCities() []models.City {
	return s.repo.GetAll()
}

func (s *cityService) GetCityByID(id int) (models.City, error) {
	return s.repo.GetByID(id)
}

func (s *cityService) CreateCity(city models.City) (models.City, error) {
	return s.repo.Create(city)
}

func (s *cityService) UpdateCity(id int, city models.City) (models.City, error) {
	return s.repo.Update(id, city)
}

func (s *cityService) DeleteCity(id int) error {
	return s.repo.Delete(id)
}
