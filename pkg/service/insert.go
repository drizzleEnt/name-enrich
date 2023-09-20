package service

import (
	nameenrich "name-enrich"
	"name-enrich/pkg/repository"
)

type InsertSrv struct {
	repo repository.Insertion
}

func NewInsertSevice(repo repository.Insertion) *InsertSrv {
	return &InsertSrv{
		repo: repo,
	}
}

func (s *InsertSrv) CreatePerson(p nameenrich.Person) (int, error) {
	return s.repo.CreatePerson(p)
}
