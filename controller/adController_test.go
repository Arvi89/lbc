package controller

import (
	"bytes"
	"encoding/json"
	"github.com/Arvi89/lbc/repository"
	"github.com/Arvi89/lbc/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAd_Post(t *testing.T) {
	carRepository := repository.NewCarMockRepository()
	carService := service.NewCarService(carRepository)

	adRepository := repository.NewAdMockRepository()
	adService := service.NewAdService(adRepository)

	adController := NewAdController(carService, adService)

	app := gin.Default()
	app.POST("/ads", adController.Post)

	w := httptest.NewRecorder()

	input, _ := json.Marshal(struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		CarModel string `json:"car_model"`
	}{
		Title:    "Test title 1",
		Content:  "Test content 1",
		CarModel: "ds 3 crossback",
	})

	req, _ := http.NewRequest("POST", "/ads", bytes.NewBuffer(input))
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAd_Get(t *testing.T) {
	carRepository := repository.NewCarMockRepository()
	carService := service.NewCarService(carRepository)

	adRepository := repository.NewAdMockRepository()
	adService := service.NewAdService(adRepository)

	adController := NewAdController(carService, adService)

	app := gin.Default()
	app.GET("/ads/:id", adController.Get)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/ads/8989", nil)
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	w = httptest.NewRecorder()

	req, _ = http.NewRequest("GET", "/ads/1111", nil)
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestAd_Put(t *testing.T) {
	carRepository := repository.NewCarMockRepository()
	carService := service.NewCarService(carRepository)

	adRepository := repository.NewAdMockRepository()
	adService := service.NewAdService(adRepository)

	adController := NewAdController(carService, adService)

	app := gin.Default()
	app.PUT("/ads/:id", adController.Put)

	w := httptest.NewRecorder()

	input, _ := json.Marshal(struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		CarModel string `json:"car_model"`
	}{
		Title:    "Test title 1",
		Content:  "Test content 1",
		CarModel: "Ds 4",
	})

	req, _ := http.NewRequest("PUT", "/ads/8989", bytes.NewBuffer(input))
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAd_Delete(t *testing.T) {
	carRepository := repository.NewCarMockRepository()
	carService := service.NewCarService(carRepository)

	adRepository := repository.NewAdMockRepository()
	adService := service.NewAdService(adRepository)

	adController := NewAdController(carService, adService)

	app := gin.Default()
	app.DELETE("/ads/:id", adController.Delete)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/ads/8989", nil)
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
