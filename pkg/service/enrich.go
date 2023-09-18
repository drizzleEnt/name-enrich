package service

import (
	"encoding/json"
	"fmt"
	nameenrich "name-enrich"
	"name-enrich/pkg/repository"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type EnrichService struct {
	repo repository.Repository
}

func NewEnrichSevice(repo repository.Repository) *EnrichService {
	return &EnrichService{
		repo: repo,
	}
}

func (e *EnrichService) EnrichAge(p *nameenrich.Person) error {

	urlAge := fmt.Sprintf("%s%s", os.Getenv("URL_AGE"), p.Name)

	resp, err := http.Get(urlAge)

	if err != nil {
		return err
	}

	err = e.decodeResponse(resp, p)

	if err != nil {
		return err
	}

	return nil
}

func (e *EnrichService) EnrichGender(p *nameenrich.Person) error {

	urlGender := fmt.Sprintf("%s%s", os.Getenv("URL_GENDER"), p.Name)

	resp, err := http.Get(urlGender)

	if err != nil {
		return err
	}

	err = e.decodeResponse(resp, p)

	if err != nil {
		return err
	}

	return nil
}

func (e *EnrichService) EnrichNationality(p *nameenrich.Person) error {

	urlNationality := fmt.Sprintf("%s%s", os.Getenv("URL_NATIONALITY"), p.Name)

	resp, err := http.Get(urlNationality)

	if err != nil {
		return err
	}

	var c nameenrich.Country

	err = json.NewDecoder(resp.Body).Decode(&c)

	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	p.Country = c.Country[0].CountryId
	return nil
}

func (e *EnrichService) decodeResponse(resp *http.Response, p *nameenrich.Person) error {
	return json.NewDecoder(resp.Body).Decode(p)
}
