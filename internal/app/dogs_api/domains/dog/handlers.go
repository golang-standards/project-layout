package dog

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	s *Service
}

func InitDogHandlers(router *gin.Engine, db *gorm.DB) {
	s := New(db)
	h := Handler{s: s}

	router.GET("/dogs", h.GetDogsHandler)
	router.GET("/dogs/:id", h.GetDogByIdHandler)
	router.POST("/dogs", h.CreateDogHandler)
	router.DELETE("/dogs/:id", h.DeleteDogHandler)
	router.PATCH("/dogs/:id", h.UpdateDogHandler)
}

func (h *Handler) GetDogsHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, h.s.GetDogs())
}

func (h *Handler) GetDogByIdHandler(c *gin.Context) {
	var id = c.Param("id")
	var dog, err = h.s.GetDogById(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.IndentedJSON(http.StatusOK, dog)
}

func (h *Handler) CreateDogHandler(c *gin.Context) {
	var dog *Dog
	if err := c.BindJSON(&dog); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	h.s.CreateDog(dog)
	c.IndentedJSON(http.StatusCreated, &dog)
}

func (h *Handler) DeleteDogHandler(c *gin.Context) {
	id := c.Param("id")
	h.s.DeleteDog(id)
}

func (h *Handler) UpdateDogHandler(c *gin.Context) {
	id := c.Param("id")

	var updateDto *UpdateDto
	if err := c.BindJSON(updateDto); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	result, err := h.s.UpdateDog(id, updateDto)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, &result)
}
