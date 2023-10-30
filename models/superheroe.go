package models

import "github.com/google/uuid"

type Superheroe struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Biography Biography `json:"biography"`
	Work      Work      `json:"work"`
}

type Biography struct {
	AlterEgo  string `json:"alter-ego"`
	FirstApp  string `json:"first-appearance"`
	Publisher string `json:"publisher"`
}

type Work struct {
	Ocupation string `json:"ocupation"`
	Base      string `json:"base"`
}
