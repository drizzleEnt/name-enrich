package service

import (
	nameenrich "name-enrich"
	"name-enrich/pkg/repository"
)

type Enrich interface {
	EnrichAge(p *nameenrich.Person) error
	EnrichGender(p *nameenrich.Person) error
	EnrichNationality(p *nameenrich.Person) error
}

type InsertService interface {
	CreatePerson(p nameenrich.Person) (int, error)
}

type ReceiptService interface {
	ReceiptPerson(id int) (nameenrich.Person, error)
}

type Service struct {
	Enrich
	InsertService
	ReceiptService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Enrich:         NewEnrichSevice(*repository),
		InsertService:  NewInsertSevice(repository.Insertion),
		ReceiptService: NewReciptService(repository.Receiption),
	}
}
