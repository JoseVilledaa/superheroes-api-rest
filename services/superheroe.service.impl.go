package services

import (
	"context"
	"errors"

	"github.com/JoseVilledaa/superheroes-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SuperheroeServiceImpl struct {
	superheroecollection *mongo.Collection
	ctx                  context.Context
}

func (s *SuperheroeServiceImpl) CreateSuperheroe(sh *models.Superheroe) error {
	_, err := s.superheroecollection.InsertOne(s.ctx, sh)
	if err != nil {
		return err
	}
	return nil
}

func NewSuperheroeService(superheroecollection *mongo.Collection, ctx context.Context) *SuperheroeServiceImpl {
	return &SuperheroeServiceImpl{
		superheroecollection: superheroecollection,
		ctx:                  ctx,
	}
}

func (s *SuperheroeServiceImpl) GetAll() ([]models.Superheroe, error) {
	var superheroe []models.Superheroe
	cursor, err := s.superheroecollection.Find(s.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(s.ctx) {
		var sh models.Superheroe
		if err = cursor.Decode(&sh); err != nil {
			return nil, err
		}
		superheroe = append(superheroe, sh)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(s.ctx)

	if len(superheroe) == 0 {
		return nil, errors.New("no superheroes yet :c")
	}
	return superheroe, nil
}
