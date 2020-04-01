package controller

import "github.com/gin-gonic/gin"

var (
	UserController controller = &userController{}
)

type controller interface {
	CreateRoutes(r *gin.Engine)
}
