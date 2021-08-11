package repository

import "github.com/Arvi89/lbc/entity"

type Ad interface {
	Get(id int) *entity.Ad
	Post(ad *entity.Ad)
	Put(ad *entity.Ad)
	Delete(id int)
}

