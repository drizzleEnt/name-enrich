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

type Authsevice interface {
	CreatePerson(p nameenrich.Person) (int, error)
}

type Service struct {
	Enrich
	Authsevice
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Enrich:     NewEnrichSevice(*repository),
		Authsevice: NewAuthSevice(repository.Autorization),
	}
}
