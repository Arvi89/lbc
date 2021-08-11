package repository

import (
	"github.com/Arvi89/lbc/entity"
	"strings"
)

type CarMock struct {
}

func NewCarMockRepository() Car {
	repository := CarMock{}

	return &repository
}

func (repository CarMock) Get(car *entity.Car) *entity.Car {
	panic("implement me")
}

func (repository CarMock) Post(car *entity.Car) {
	panic("implement me")
}

func (repository CarMock) Put(car *entity.Car) {
	panic("implement me")
}

func (repository CarMock) Delete(id int) {
	panic("implement me")
}

func (repository CarMock) FindByName(name string) (cars []*entity.Car) {
	citroen := entity.Brand{
		ID:   3,
		Name: "Citroen",
		Cars: nil,
	}

	audi := entity.Brand{
		ID:   1,
		Name: "Audi",
		Cars: nil,
	}

	bmw := entity.Brand{
		ID:   2,
		Name: "BMW",
		Cars: nil,
	}

	switch strings.ToLower(name) {
	case "ds":
		cars = append(cars, &entity.Car{
			ID:          2,
			Brand:       3,
			BrandStruct: citroen,
			Model:       "Ds4",
		}, &entity.Car{
			ID:          1,
			Brand:       3,
			BrandStruct: citroen,
			Model:       "Ds3",
		})
	case "3":
		cars = append(cars, &entity.Car{
			ID:          4,
			Brand:       1,
			BrandStruct: audi,
			Model:       "Q3",
		}, &entity.Car{
			ID:          5,
			Brand:       3,
			BrandStruct: citroen,
			Model:       "Rs3",
		}, &entity.Car{
			ID:          6,
			Brand:       2,
			BrandStruct: bmw,
			Model:       "Serie 3",
		}, &entity.Car{
			ID:          7,
			Brand:       3,
			BrandStruct: citroen,
			Model:       "Ds3",
		})
	case "4":
		cars = append(cars, &entity.Car{
			ID:          8,
			Brand:       1,
			BrandStruct: audi,
			Model:       "Rs4",
		}, &entity.Car{
			ID:          9,
			Brand:       1,
			BrandStruct: audi,
			Model:       "S4",
		}, &entity.Car{
			ID:          10,
			Brand:       1,
			BrandStruct: audi,
			Model:       "S4 Avant",
		}, &entity.Car{
			ID:          11,
			Brand:       2,
			BrandStruct: bmw,
			Model:       "Serie 4",
		}, &entity.Car{
			ID:          12,
			Brand:       3,
			BrandStruct: citroen,
			Model:       "Ds4",
		})
	}

	return cars
}
