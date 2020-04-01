package controller

import (
	"net/http"
	"strconv"

	"github.com/Dylvn/dragonfire-api/src/models/users"
	"github.com/Dylvn/dragonfire-api/src/services"
	"github.com/gin-gonic/gin"
)

type userController struct{}

func (uc *userController) CreateRoutes(r *gin.Engine) {
	r.POST("/users", uc.Store)
	r.GET("/users/:userID", uc.Show)
	r.PUT("/users/:userID", uc.Update)
	r.DELETE("/users/:userID", uc.Delete)
}

func (uc *userController) Store(c *gin.Context) {
	var u users.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid json",
		})
		return
	}

	user, err := services.UsersService.StoreUser(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (uc *userController) Show(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	u, err := services.UsersService.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, u)
}

func (uc *userController) Update(c *gin.Context) {
	var u users.User

	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid json",
		})
		return
	}

	u.Id = userID

	user, err := services.UsersService.UpdateUser(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *userController) Delete(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if err := services.UsersService.DeleteUser(users.User{Id: userID}); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"status": "deleted",
	})
}
