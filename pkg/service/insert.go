package service

import (
	nameenrich "name-enrich"
	"name-enrich/pkg/repository"
)

type AuthService struct {
	repo repository.Autorization
}

func NewAuthSevice(repo repository.Autorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreatePerson(p nameenrich.Person) (int, error) {
	return s.repo.CreatePerson(p)
}
