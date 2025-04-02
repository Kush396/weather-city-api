package repository

import (
	"errors"

	"weather-city-api/models"
)

type CityRepository interface {
	GetAll() []models.City
	GetByID(id int) (models.City, error)
	Create(city models.City) (models.City, error)
	Update(id int, city models.City) (models.City, error)
	Delete(id int) error
}

type cityRepository struct {
	cities map[int]models.City
	nextID int
}

func NewCityRepository() CityRepository {
	return &cityRepository{
		cities: make(map[int]models.City),
		nextID: 1,
	}
}

func (r *cityRepository) GetAll() []models.City {
	cities := make([]models.City, 0, len(r.cities))
	for _, city := range r.cities {
		cities = append(cities, city)
	}
	return cities
}

func (r *cityRepository) GetByID(id int) (models.City, error) {
	city, exists := r.cities[id]
	if !exists {
		return models.City{}, errors.New("city not found")
	}
	return city, nil
}

func (r *cityRepository) Create(city models.City) (models.City, error) {
	city.ID = r.nextID
	r.cities[city.ID] = city
	r.nextID++

	return city, nil
}

func (r *cityRepository) Update(id int, city models.City) (models.City, error) {
	exists := false
	_, exists = r.cities[id]
	if !exists {
		return models.City{}, errors.New("city not found")
	}

	city.ID = id
	r.cities[id] = city

	return city, nil
}

func (r *cityRepository) Delete(id int) error {
	exists := false
	_, exists = r.cities[id]
	if !exists {
		return errors.New("city not found")
	}

	delete(r.cities, id)
	return nil
}
