package repository

import (
	"fmt"
	nameenrich "name-enrich"

	"github.com/jmoiron/sqlx"
)

const (
	personTable = "users"
)

type Insertpostgres struct {
	db *sqlx.DB
}

func NewInsertpostgres(db *sqlx.DB) *Insertpostgres {
	return &Insertpostgres{
		db: db,
	}
}

func (a *Insertpostgres) CreatePerson(p nameenrich.Person) (int, error) {

	tx, err := a.db.Begin()

	if err != nil {
		return 0, err
	}

	var id int

	createItemQuary := fmt.Sprintf("INSERT INTO %s (name, surname, patronymic, age, gender, nationality) values ($1, $2, $3, $4, $5, $6) RETURNING id", personTable)
	row := tx.QueryRow(createItemQuary, p.Name, p.Surname, p.Patronymic, p.Age, p.Gender, p.Nationality)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
