package repository

import (
	"github.com/Arvi89/lbc/entity"
	"math/rand"
)

type AdMock struct {
}

func NewAdMockRepository() Ad {
	repository := AdMock{}

	return repository
}

func (repository AdMock) Get(id int) *entity.Ad {
	switch id {
	case 8989:
		d := 12
		return &entity.Ad{
			Id:      8989,
			Title:   "Test ad 8989",
			Content: "Content 8989",
			Car:     &d,
			CarStruct: &entity.Car{
				ID:    12,
				Brand: 3,
				BrandStruct: entity.Brand{
					ID:   3,
					Name: "Citroen",
					Cars: nil,
				},
				Model: "Super Car",
			},
		}
	case 9090:
		d := 13
		return &entity.Ad{
			Id:      9090,
			Title:   "Test ad 9090",
			Content: "Content 9090",
			Car:     &d,
			CarStruct: &entity.Car{
				ID:    15,
				Brand: 2,
				BrandStruct: entity.Brand{
					ID:   2,
					Name: "BMW",
					Cars: nil,
				},
				Model: "Super Car 2",
			},
		}
	}

	return nil
}

func (repository AdMock) Post(ad *entity.Ad) {
	ad.Id = rand.Int()
}

func (repository AdMock) Put(ad *entity.Ad) {
}

func (repository AdMock) Delete(id int) {
}
