package handler

import (
	"l0s0s/Rarible-API/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetNFTOwnership(id string) (model.Ownership, error)
	GetNFTTraitsRarity(collectionID string, properties []model.TraitProperty) (model.TraitsRarity, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetNFTOwnership(c *gin.Context) {
	id := c.Param("id")

	ownership, err := h.service.GetNFTOwnership(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ownership)
}
func (h *Handler) GetNFTTraitsRarity(c *gin.Context) {
	collectionID := c.Param("collectionID")

	var properties []model.TraitProperty
	if err := c.BindJSON(&properties); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid properties format"})
		return
	}

	traitsRarity, err := h.service.GetNFTTraitsRarity(collectionID, properties)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, traitsRarity)
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.GET("/nft/ownership/:id", h.GetNFTOwnership)
	router.POST("/nft/traits/rarity/:collectionID", h.GetNFTTraitsRarity)
}
