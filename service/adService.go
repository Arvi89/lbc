package service

import (
	"github.com/Arvi89/lbc/entity"
	"github.com/Arvi89/lbc/repository"
)

type Ad struct {
	adRepository repository.Ad
}

func NewAdService(adRepository repository.Ad) *Ad {
	service := Ad{
		adRepository: adRepository,
	}

	return &service
}

func (service Ad) GetAd(id int) *entity.Ad {
	ad := service.adRepository.Get(id)

	return ad
}

func (service Ad) DeleteAd(id int) error {
	service.adRepository.Delete(id)

	return nil
}

func (service Ad) PostCarAd(title string, content string, car *entity.Car) *entity.Ad {
	ad := entity.Ad{
		Title:   title,
		Content: content,
		Car:     &car.ID,
	}

	service.adRepository.Post(&ad)

	if ad.Id != 0 {
		return &ad
	}

	return nil
}

func (service Ad) PutAd(ad *entity.Ad) {
	service.adRepository.Put(ad)
}
