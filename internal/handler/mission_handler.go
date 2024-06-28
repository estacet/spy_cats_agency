package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"spy-cats/internal/service"
)

type MissionCRUDHandler struct {
	missionService *service.MissionService
}

func NewMissionCRUDHandler(missionService *service.MissionService) *MissionCRUDHandler {
	return &MissionCRUDHandler{
		missionService: missionService,
	}
}

func (h *MissionCRUDHandler) RegisterRoutes(router *gin.Engine) {
	const basePath = "/mission"
	const resourcePath = basePath + "/:id"

	router.GET(resourcePath, h.get)
	router.GET(basePath, h.getList)
	router.PATCH(resourcePath, h.update)
	router.DELETE(resourcePath, h.delete)
}

func (h *MissionCRUDHandler) get(c *gin.Context) {
	id := c.Params.ByName("id")

	parsedId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get mission by id",
			"error":   err.Error(),
		})

		return
	}

	mission, err := h.missionService.GetById(c, parsedId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get mission by id",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"mission": mission})
}

func (h *MissionCRUDHandler) getList(c *gin.Context) {
	missionList, err := h.missionService.GetList(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get mission list",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"missions": missionList})
}

func (h *MissionCRUDHandler) update(c *gin.Context) {
	id := c.Params.ByName("id")

	parsedId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete mission by id",
			"error":   err.Error(),
		})

		return
	}

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update mission profile",
			"error":   err.Error(),
		})

		return
	}
	defer c.Request.Body.Close()

	args := new(service.UpdateMissingArgs)

	if err := json.Unmarshal(bodyBytes, args); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update mission profile",
			"error":   err.Error(),
		})

		return
	}

	err = h.missionService.Update(c, parsedId, args)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update mission by id",
			"error":   err.Error(),
		})

		return
	}

	c.Status(http.StatusOK)
}

func (h *MissionCRUDHandler) delete(c *gin.Context) {
	id := c.Params.ByName("id")

	parsedId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete mission by id",
			"error":   err.Error(),
		})

		return
	}

	err = h.missionService.Delete(c, parsedId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete mission by id",
			"error":   err.Error(),
		})

		return
	}

	c.Status(http.StatusNoContent)
}
