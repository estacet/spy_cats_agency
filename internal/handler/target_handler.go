package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"spy-cats/internal/service"
)

type TargetCRUDHandler struct {
	targetService *service.TargetService
}

func NewTargetCRUDHandler(targetService *service.TargetService) *TargetCRUDHandler {
	return &TargetCRUDHandler{targetService: targetService}
}

func (h *TargetCRUDHandler) RegisterRoutes(router *gin.Engine) {
	const basePath = "/mission_target"
	const resourcePath = basePath + "/:id"

	router.POST(basePath, h.create)
	router.PATCH(resourcePath, h.update)
}

func (h *TargetCRUDHandler) create(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to add target",
			"error":   err.Error(),
		})

		return
	}

	defer c.Request.Body.Close()

	args := new(service.CreateTargetArgs)

	if err := json.Unmarshal(bodyBytes, args); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to add target",
			"error":   err.Error(),
		})

		return
	}

	id, err := h.targetService.Create(c, args)
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *TargetCRUDHandler) update(c *gin.Context) {
	id := c.Params.ByName("id")

	bodyBytes, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update target",
			"error":   err.Error(),
		})

		return
	}

	defer c.Request.Body.Close()

	args := new(service.UpdateTargetArgs)

	if err := json.Unmarshal(bodyBytes, args); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update target",
			"error":   err.Error(),
		})

		return
	}

	parsedId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update target",
			"error":   err.Error(),
		})

		return
	}

	err = h.targetService.Update(c, parsedId, args)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update target",
			"error":   err.Error(),
		})

		return
	}

	c.Status(http.StatusOK)
}
