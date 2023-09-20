package repository

import (
	nameenrich "name-enrich"

	"github.com/jmoiron/sqlx"
)

type Autorization interface {
	CreatePerson(p nameenrich.Person) (int, error)
}

type Repository struct {
	Autorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: NewAuthpostgres(db),
	}
}
