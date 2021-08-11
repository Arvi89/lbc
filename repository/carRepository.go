package repository

import "github.com/Arvi89/lbc/entity"

type Car interface {
	Get(car *entity.Car) *entity.Car
	Post(car *entity.Car)
	Put(car *entity.Car)
	Delete(id int)
	FindByName(name string) []*entity.Car
}
