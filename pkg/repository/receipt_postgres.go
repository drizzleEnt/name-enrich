package repository

import (
	"fmt"
	nameenrich "name-enrich"

	"github.com/jmoiron/sqlx"
)

type ReceiptPostgres struct {
	db *sqlx.DB
}

func NewReceiptpostgres(db *sqlx.DB) *ReceiptPostgres {
	return &ReceiptPostgres{
		db: db,
	}
}

func (r *ReceiptPostgres) ReceiptPerson(id int) (nameenrich.Person, error) {

	var p nameenrich.Person

	//TODO: Check cash

	quary := fmt.Sprintf("SELECT * FROM %s WHERE \"id\" = $1", personTable)
	err := r.db.Get(&p, quary, id)

	return p, err
}
