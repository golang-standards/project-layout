package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	service *Service
}

func InitUserHandlers(router *gin.Engine, db *gorm.DB) {
	handler := Handler{service: InitUserService(db)}

	router.GET("/users", handler.GetAllHandler)
	router.GET("/users/:id", handler.GetByIdHandler)
	router.POST("/users", handler.CreateHandler)
	router.DELETE("/users/:id", handler.DeleteHandler)
	router.PATCH("/users/:id", handler.UpdateHandler)
}

func (h *Handler) GetAllHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, h.service.GetAll())
}

func (h *Handler) GetByIdHandler(c *gin.Context) {
	var id = c.Param("id")
	var user, err = h.service.GetById(id)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (h *Handler) CreateHandler(c *gin.Context) {
	var user *User
	if err := c.BindJSON(&user); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	h.service.Create(user)
	c.IndentedJSON(http.StatusCreated, &user)
}

func (h *Handler) DeleteHandler(c *gin.Context) {
	id := c.Param("id")
	h.service.Delete(id)
}

func (h *Handler) UpdateHandler(c *gin.Context) {
	id := c.Param("id")

	var updateDto *UpdateDto
	if err := c.BindJSON(updateDto); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	result, err := h.service.Update(id, updateDto)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, &result)
}
