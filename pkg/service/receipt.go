package service

import (
	nameenrich "name-enrich"
	"name-enrich/pkg/repository"
)

type ReceiptSrv struct {
	repo repository.Receiption
}

func NewReciptService(repo repository.Receiption) *ReceiptSrv {
	return &ReceiptSrv{
		repo: repo,
	}
}

func (s *ReceiptSrv) ReceiptPerson(id int) (nameenrich.Person, error) {
	return s.repo.ReceiptPerson(id)
}
