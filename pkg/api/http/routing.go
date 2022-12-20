package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoutesConfigurer struct {
	UserController UserController
}

func (rc RoutesConfigurer) ConfigureRoutes(engine *gin.Engine) {
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	userGrp := engine.Group("/user")
	userGrp.GET("", rc.UserController.GetUsersByQuery)
	userGrp.PATCH("", rc.UserController.UpdateUser)
	userGrp.PUT("", rc.UserController.CreateUser)
	userGrp.DELETE("", rc.UserController.DeleteUser)
}
