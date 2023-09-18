package handler

import (
	nameenrich "name-enrich"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) checkPerson(c *gin.Context) {
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

	h.EnrichInformation(c, &p)
}

func (h *Handler) EnrichInformation(c *gin.Context, p *nameenrich.Person) {
	err := h.service.Enrich.EnrichAge(p)

	if err != nil {
		logrus.Error("faild to connect api " + err.Error())
		errorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	err = h.service.EnrichGender(p)

	if err != nil {
		errorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	err = h.service.EnrichNationality(p)

	if err != nil {
		errorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"person": p,
	})

}
