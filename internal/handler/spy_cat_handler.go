package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"io"
	"net/http"
	"spy-cats/internal/service"
)

type SpyCatCRUDHandler struct {
	spyCatService *service.SpyCatService
	validate      *validator.Validate
}

func NewSpyCatCRUDHandler(
	spyCatService *service.SpyCatService,
	validate *validator.Validate,
) *SpyCatCRUDHandler {
	return &SpyCatCRUDHandler{
		spyCatService: spyCatService,
		validate:      validate,
	}
}

func (h *SpyCatCRUDHandler) RegisterRoutes(router *gin.Engine) {
	const basePath = "/spyCat"
	const resourcePath = basePath + "/:id"

	router.POST(basePath, h.create)
	router.GET(resourcePath, h.get)
	router.GET(basePath, h.getList)
	router.PATCH(resourcePath, h.update)
}

func (h *SpyCatCRUDHandler) create(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create spy cat profile",
			"error":   err.Error(),
		})

		return
	}

	defer c.Request.Body.Close()

	args := new(service.CreateSpyCatArgs)

	if err := json.Unmarshal(bodyBytes, args); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create spy cat profile",
			"error":   err.Error(),
		})

		return
	}

	err = h.validate.Struct(args)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create spy cat profile",
			"error":   err.Error(),
		})

		return
	}

	id, err := h.spyCatService.Create(c, args)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create spy cat profile",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *SpyCatCRUDHandler) get(c *gin.Context) {
	id := c.Params.ByName("id")

	parsedId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get spy cat by id",
			"error":   err.Error(),
		})

		return
	}

	spyCat, err := h.spyCatService.GetById(c, parsedId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get spy cat by id",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"spy_cat": spyCat})
}

func (h *SpyCatCRUDHandler) update(c *gin.Context) {
	id := c.Params.ByName("id")

	bodyBytes, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update spy cat profile",
			"error":   err.Error(),
		})

		return
	}

	defer c.Request.Body.Close()

	args := new(service.UpdateSpyCatArgs)

	if err := json.Unmarshal(bodyBytes, args); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update spy cat profile",
			"error":   err.Error(),
		})

		return
	}

	parsedId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update spy cat profile",
			"error":   err.Error(),
		})

		return
	}

	err = h.spyCatService.Update(c, parsedId, args)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update spy cat profile",
			"error":   err.Error(),
		})

		return
	}

	c.Status(http.StatusOK)
}

func (h *SpyCatCRUDHandler) getList(c *gin.Context) {
	spyCatsList, err := h.spyCatService.GetList(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get spy cats list",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"spy_cats": spyCatsList})
}
