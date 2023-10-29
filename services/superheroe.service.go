package services

import "github.com/JoseVilledaa/superheroes-api/models"

type SuperheroeService interface {
	GetAll() ([]models.Superheroe, error)
}
