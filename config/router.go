package config

import (
	"github.com/gin-gonic/gin"
	"github.com/Arvi89/lbc/controller"
)

type Router struct {
	gin           *gin.Engine
	adController  *controller.Ad
}

func NewRouting(gin *gin.Engine, adController *controller.Ad) *Router {
	router := Router{
		gin:           gin,
		adController: adController,
	}
	router.initRoutes()

	return &router
}

func (router Router) initRoutes() {
	router.gin.POST("/ads", router.adController.Post)
	router.gin.GET("/ads/:id", router.adController.Get)
	router.gin.PUT("/ads/:id", router.adController.Put)
	router.gin.DELETE("/ads/:id", router.adController.Delete)
}
