package service

import (
	nameenrich "name-enrich"
	"name-enrich/pkg/repository"
	"net/http"
)

type Enrich interface {
	EnrichAge(p *nameenrich.Person) (*http.Response, error)
	EnrichGender(p *nameenrich.Person) (*http.Response, error)
	EnrichNationality(p *nameenrich.Person) (*http.Response, error)
}

type Service struct {
	Enrich
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Enrich: NewEnrichSevice(*repository),
	}
}
