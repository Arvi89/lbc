package repository

import (
	"gorm.io/gorm"
	"github.com/Arvi89/lbc/entity"
)

type AdMysql struct {
	db *gorm.DB
}

func NewAdMysqlRepository(db *gorm.DB) Ad {
	repository := AdMysql{
		db: db,
	}

	return &repository
}

func (repository AdMysql) Get(id int) (ad *entity.Ad) {
	repository.db.
		Preload("CarStruct").
		Preload("CarStruct.BrandStruct").
		Where("ID = ?", id).
		First(&ad)

	return ad
}

func (repository AdMysql) Post(ad *entity.Ad) {
	repository.db.Create(ad)
}

func (repository AdMysql) Put(ad *entity.Ad) {
	repository.db.Save(ad)
}

func (repository AdMysql) Delete(id int) {
	repository.db.Delete(&entity.Ad{}, id)
}
