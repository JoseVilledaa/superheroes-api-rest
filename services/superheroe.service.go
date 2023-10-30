package services

import (
	"github.com/JoseVilledaa/superheroes-api/models"
	"github.com/google/uuid"
)

type SuperheroeService interface {
	CreateSuperheroe(*models.Superheroe) error
	GetAll() ([]models.Superheroe, error)
	DeleteSuperheroe(id uuid.UUID) error
}
