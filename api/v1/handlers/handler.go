package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
)

type service interface {
	GetAuction(ctx context.Context, id string) (map[string]interface{}, error)
	GetProduct(ctx context.Context, id string) (map[string]interface{}, error)
	Search(ctx context.Context, req string, page int) ([]interface{}, error)
}

type Handler struct {
	service service
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "service works..."})
}

func NewHandler(s service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) GetAuction(c *gin.Context) {
	id := c.Param("id")
	data, err := h.service.GetAuction(context.TODO(), id)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	c.JSON(200, data)
}

func (h *Handler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	data, err := h.service.GetProduct(context.TODO(), id)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	c.JSON(200, data)
}

func (h *Handler) Suggestions(c *gin.Context) {
	
}
