package models

type City struct {
	ID          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	Country     string `json:"country"`
	Description string `json:"description,omitempty"`
}

type CityRequest struct {
	Name        string `json:"name" binding:"required"`
	Country     string `json:"country"`
	Description string `json:"description,omitempty"`
}
