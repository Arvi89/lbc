package service

import (
	"github.com/Arvi89/lbc/repository"
	"testing"
)

func TestCarService_match(t *testing.T) {
	carRepository := repository.NewCarMockRepository()
	carService := NewCarService(carRepository)
	car := carService.Match("ds 3 crossback")

	if car == nil {
		t.Errorf("Expected Ds3, got nothing")
	}

	if car.Model != "Ds3" {
		if car == nil {
			t.Errorf("Expected Ds3, got %s", car.Model)
		}
	}
}
