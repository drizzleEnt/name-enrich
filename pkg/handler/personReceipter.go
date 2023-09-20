package handler

import (
	nameenrich "name-enrich"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ReceiptPersonFromDb(c *gin.Context) {
	var p nameenrich.Person

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	p, err = h.service.ReceiptService.ReceiptPerson(id)

	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"person": p,
	})
}
