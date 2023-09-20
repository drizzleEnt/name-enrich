package handler

import (
	nameenrich "name-enrich"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) takePerson(c *gin.Context) {
	var p nameenrich.Person

	if err := c.BindJSON(&p); err != nil {
		logrus.Error(err.Error())
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if p.Name == "" || p.Surname == "" {
		logrus.Error("incorrect data ")
		errorResponse(c, http.StatusBadRequest, "name or surname is required")
		return
	}

	err := h.EnrichInformation(c, &p)

	if err != nil {
		errorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	id, err := h.service.Authsevice.CreatePerson(p)

	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
	}

	p.Id = id

	c.JSON(http.StatusOK, map[string]interface{}{
		"person": p,
	})
}

func (h *Handler) EnrichInformation(c *gin.Context, p *nameenrich.Person) error {
	err := h.service.Enrich.EnrichAge(p)

	if err != nil {
		return err
	}

	err = h.service.EnrichGender(p)

	if err != nil {
		return err
	}

	err = h.service.EnrichNationality(p)

	if err != nil {
		return err
	}

	return nil
}
