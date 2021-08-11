package repository

import (
	"gorm.io/gorm"
	"github.com/Arvi89/lbc/entity"
)

type CarMysql struct {
	db *gorm.DB
}

func NewCarMysqlRepository(db *gorm.DB) Car {
	repository := CarMysql{
		db: db,
	}

	return &repository
}

func (repository CarMysql) Get(car *entity.Car) *entity.Car {
	repository.db.First(car)

	return car
}

func (repository CarMysql) Post(car *entity.Car) {
	repository.db.Create(car)
}

func (repository CarMysql) Put(car *entity.Car) {
	repository.db.Save(car)
}

func (repository CarMysql) Delete(id int) {
	repository.db.Delete(&entity.Car{}, id)
}

func (repository CarMysql) FindByName(name string) (cars []*entity.Car) {
	repository.db.
		Preload("BrandStruct").
		Where("model LIKE ?", "%"+name+"%").
		Find(&cars)

	return cars
}
