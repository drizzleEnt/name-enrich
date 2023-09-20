package handler

import (
	"name-enrich/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	getPerson := router.Group("/")
	{
		getPerson.POST("/", h.takePerson)
		getPerson.GET("/:id", h.ReceiptPersonFromDb)
	}

	return router
}
