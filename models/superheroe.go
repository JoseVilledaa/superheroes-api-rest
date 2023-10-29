package models

type Superheroe struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Biography Biography `json:"biography"`
}

type Biography struct {
	FullName string `json:"full-name"`
	AlterEgo string `json:"alter-ego"`
	FirstApp string `json:"first-appearance"`
}
