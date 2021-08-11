package service

import (
	"github.com/Arvi89/lbc/entity"
	"github.com/Arvi89/lbc/repository"
	"strings"
)

type Car struct {
	carRepository repository.Car
}

func NewCarService(carRepository repository.Car) *Car {
	service := Car{
		carRepository: carRepository,
	}

	return &service
}

func (service Car) Match(input string) *entity.Car {
	var cars [][]*entity.Car
	var carToReturn *entity.Car
	s := strings.Split(input, " ")

	for _, part := range s {
		temp := service.carRepository.FindByName(part)
		if len(temp) > 0 {
			cars = append(cars, temp)
		}
	}

	l := len(cars)
	if l > 1 {
		var backupCar *entity.Car
		for key, arr := range cars {
			for _, car := range arr {
				if strings.ToLower(car.Model) == strings.ToLower(s[key]) {
					backupCar = car
				}
				
				for i := key+1; i < l; i++ {
					for _, car2 := range cars[i] {
						if car.Model == car2.Model {
							// If we get more than one Match, we don't know which one is right.
							// Ideally, should return a proper error
							if carToReturn != nil {
								return nil
							}

							carToReturn = car
						}
					}
				}
			}
		}

		if backupCar != nil && carToReturn == nil {
			return backupCar
		}
	} else if len(cars) == 1 {
		for _, car := range cars[0] {
			for _, part := range s {
				if strings.ToLower(car.Model) == strings.ToLower(part) {
					return car
				}
			}
		}

		// If we have retrieved more than 1 car with only one word, we can't know which car to return
		if len(cars[0]) != 1 {
			return nil
		}

		carToReturn = cars[0][0]
	}

	return carToReturn
}
