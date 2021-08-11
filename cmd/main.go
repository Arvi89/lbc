package main

import (
	"github.com/Arvi89/lbc/config"
	"github.com/Arvi89/lbc/controller"
	"github.com/Arvi89/lbc/repository"
	"github.com/Arvi89/lbc/service"
	"github.com/gin-gonic/gin"
)

func main() {
	appConfig := config.NewApp()

	carRepository := repository.NewCarMysqlRepository(appConfig.Db)
	adRepository := repository.NewAdMysqlRepository(appConfig.Db)

	carService := service.NewCarService(carRepository)
	adService := service.NewAdService(adRepository)

	adController := controller.NewAdController(carService, adService)


	app := gin.Default()
	config.NewRouting(app, adController)

	_ = app.Run(":4000")
}
