package dog

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	service *Service
}

func InitDogHandlers(router *gin.Engine, db *gorm.DB) {
	dogService := InitDogService(db)
	handler := Handler{service: dogService}

	router.GET("/dogs", handler.GetDogsHandler)
	router.GET("/dogs/:id", handler.GetDogByIdHandler)
	router.POST("/dogs", handler.CreateDogHandler)
	router.DELETE("/dogs/:id", handler.DeleteDogHandler)
	router.PATCH("/dogs/:id", handler.UpdateDogHandler)
}

func (h *Handler) GetDogsHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, h.service.GetDogs())
}

func (h *Handler) GetDogByIdHandler(c *gin.Context) {
	var id = c.Param("id")
	var dog, err = h.service.GetDogById(id)
	if err != nil {
		if err.Error() == "record not found" {
			_ = c.AbortWithError(http.StatusNotFound, err)
			return
		} else {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
	}
	c.IndentedJSON(http.StatusOK, dog)
}

func (h *Handler) CreateDogHandler(c *gin.Context) {
	var dog *Dog
	if err := c.BindJSON(&dog); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	h.service.CreateDog(dog)
	c.IndentedJSON(http.StatusCreated, &dog)
}

func (h *Handler) DeleteDogHandler(c *gin.Context) {
	id := c.Param("id")
	h.service.DeleteDog(id)
}

func (h *Handler) UpdateDogHandler(c *gin.Context) {
	id := c.Param("id")

	var updateDto *UpdateDto
	if err := c.BindJSON(updateDto); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	result, err := h.service.UpdateDog(id, updateDto)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, &result)
}
