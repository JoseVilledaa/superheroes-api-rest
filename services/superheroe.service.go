package services

import "github.com/JoseVilledaa/superheroes-api/models"

type SuperheroeService interface {
	CreateSuperheroe(*models.Superheroe) error
	GetAll() ([]models.Superheroe, error)
}
