package app

import (
	"github.com/Dylvn/dragonfire-api/src/controller"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func Start() {
	controller.UserController.CreateRoutes(r)
	r.Run(":8080")
}
