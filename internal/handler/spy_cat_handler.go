package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"spy-cats/internal/service"
)

type SpyCatCRUDHandler struct {
	spyCatService *service.SpyCatService
}

func NewSpyCatCRUDHandler(
	spyCatService *service.SpyCatService,
) *SpyCatCRUDHandler {
	return &SpyCatCRUDHandler{
		spyCatService: spyCatService,
	}
}

func (h *SpyCatCRUDHandler) RegisterRoutes(router *gin.Engine) {
	const basePath = "/spyCat"
	const resourcePath = basePath + "/:id"

	router.POST(basePath, h.create)
	router.GET(resourcePath, h.get)
	//r.GET(basePath, h.getList(*gin.Context))
	//r.PATCH(resourcePath, h.update(*gin.Context))
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

	err = h.spyCatService.Create(c, args)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create spy cat profile",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successful Spy cat profile creation."})
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

//
//func (h *SpyCatCRUDHandler) update(c *gin.Context) {
//	id := c.Params.ByName("id")
//
//	bodyBytes, err := io.ReadAll(r.Body)
//	if err != nil {
//		server.HandleError(w, err)
//
//		return
//	}
//	defer r.Body.Close()
//
//	args := new(service.UpdateDriverArgs)
//
//	if err := json.Unmarshal(bodyBytes, args); err != nil {
//		server.HandleError(w, err)
//
//		return
//	}
//
//	err = h.validate.Struct(args)
//	if err != nil {
//		server.HandleError(w, err)
//
//		return
//	}
//
//	parsedId, err := uuid.Parse(id)
//	if err != nil {
//		server.HandleError(w, err)
//
//		return
//	}
//
//	err = h.driverService.Update(r.Context(), parsedId, args)
//
//	if err != nil {
//		server.HandleError(w, err)
//
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//}

//func (h *SpyCatCRUDHandler) getList(c *gin.Context) {
//	driversList, err := h.driverService.GetList(r.Context())
//	if err != nil {
//		server.HandleError(w, err)
//
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//
//	if err := json.NewEncoder(w).Encode(driversList); err != nil {
//		server.HandleError(w, err)
//
//		return
//	}
//}
//
