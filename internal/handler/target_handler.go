package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
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
	//router.GET(resourcePath, h.get)
	//router.GET(basePath, h.getList)
	//router.PATCH(resourcePath, h.update)
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

	err = h.targetService.Create(c, args)
}
