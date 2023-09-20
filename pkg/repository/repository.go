package repository

import (
	nameenrich "name-enrich"

	"github.com/jmoiron/sqlx"
)

type Insertion interface {
	CreatePerson(p nameenrich.Person) (int, error)
}

type Receiption interface {
	ReceiptPerson(int) (nameenrich.Person, error)
}

type Repository struct {
	Insertion
	Receiption
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Insertion:  NewInsertpostgres(db),
		Receiption: NewReceiptpostgres(db),
	}
}
