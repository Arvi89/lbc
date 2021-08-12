package controller

import (
	"github.com/Arvi89/lbc/entity"
	"github.com/Arvi89/lbc/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Ad struct {
	carService *service.Car
	adService  *service.Ad
}

type AdOutput struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Car      *struct {
		Brand string `json:"brand"`
		Model string `json:"model"`
	} `json:"car,omitempty"`
	Job     *struct{} `json:"job,omitempty"`
	Housing *struct{} `json:"housing,omitempty"`
}

type adInput struct {
	Title            string `json:"title"`
	Content          string `json:"content"`
	CarModel         string `json:"car_model"`
	JobDescription   string `json:"job_description"`
	HouseDescription string `json:"house_description"`
}

const (
	HOUSING = "housing"
	AUTO    = "auto"
	JOB     = "job"
)

func NewAdController(carService *service.Car, adService *service.Ad) *Ad {
	controller := Ad{
		carService: carService,
		adService:  adService,
	}

	return &controller
}

func (controller Ad) Get(ctx *gin.Context) {
	stringId := ctx.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{
			Message: "bad ID",
		})
		return
	}

	ad := controller.adService.GetAd(id)

	if ad == nil {
		ctx.JSON(http.StatusNotFound, struct {
			Message string `json:"message"`
		}{
			Message: "Ad not found",
		})
		return
	}

	output := AdOutput{
		Id:       ad.Id,
		Title:    ad.Title,
		Content:  ad.Content,
		Category: controller.getCategory(ad),
	}
	controller.fillOutput(&output, ad)

	ctx.JSON(200, output)
}

func (controller Ad) getCategory(ad *entity.Ad) string {
	if ad.CarStruct != nil {
		return AUTO
	}

	//TODO properly handle other cases
	return HOUSING
}

func (controller Ad) fillOutput(output *AdOutput, ad *entity.Ad) {
	switch {
	case ad.CarStruct != nil:
		output.Car = &struct {
			Brand string `json:"brand"`
			Model string `json:"model"`
		}{
			Brand: ad.CarStruct.BrandStruct.Name,
			Model: ad.CarStruct.Model,
		}
	}
	//TODO handle other cases
}

func (controller Ad) Post(ctx *gin.Context) {
	var input adInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{
			Message: "Wrong ad format",
		})
		return
	}

	if input.Title == "" || input.Content == "" {
		ctx.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{
			Message: "Wrong ad format",
		})
		return
	}

	if input.CarModel != "" {
		controller.postCarAd(ctx, input)
		return
	}

	//TODO Handle other cases
	ctx.JSON(http.StatusBadRequest, struct {
		Message string `json:"message"`
	}{
		Message: "Wrong ad format",
	})
	return
}

func (controller Ad) postCarAd(ctx *gin.Context, input adInput) {
	car := controller.carService.Match(input.CarModel)

	if car == nil {
		ctx.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{
			Message: "Can't find that car model: " + input.CarModel,
		})
		return
	}

	ad := controller.adService.PostCarAd(input.Title, input.Content, car)

	if ad == nil {
		ctx.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{
			Message: "There were error while posting the ad",
		})
		return
	}

	ctx.JSON(http.StatusOK, AdOutput{
		Id:       ad.Id,
		Title:    ad.Title,
		Content:  ad.Content,
		Category: AUTO,
		Car: &struct {
			Brand string `json:"brand"`
			Model string `json:"model"`
		}{
			Brand: car.BrandStruct.Name,
			Model: car.Model,
		},
	})
}

func (controller Ad) Put(ctx *gin.Context) {
	stringId := ctx.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{
			Message: "bad ID",
		})
		return
	}

	beforeAd := controller.adService.GetAd(id)
	if beforeAd == nil {
		ctx.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{
			Message: "This ad does not exist",
		})
		return
	}

	var input adInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{
			Message: "Wrong ad format",
		})
		return
	}

	ad := &entity.Ad{
		Id: id,
	}

	if input.Title != "" {
		ad.Title = input.Title
	}

	if input.Content != "" {
		ad.Content = input.Content
	}

	if input.CarModel != "" {
		car := controller.carService.Match(input.CarModel)
		if car == nil {
			ctx.JSON(http.StatusBadRequest, struct {
				Message string `json:"message"`
			}{
				Message: "Can't find that car model: " + input.CarModel,
			})
			return
		}
		ad.CarStruct = car
		controller.adService.PutAd(ad)

		ctx.JSON(http.StatusOK, AdOutput{
			Id:       ad.Id,
			Title:    ad.Title,
			Content:  ad.Content,
			Category: controller.getCategory(ad),
			Car: &struct {
				Brand string `json:"brand"`
				Model string `json:"model"`
			}{
				Brand: car.BrandStruct.Name,
				Model: car.Model,
			},
			Job:     nil,
			Housing: nil,
		})
		return
	}

	//TODO Handle other cases
	ctx.JSON(http.StatusBadRequest, struct {
		Message string `json:"message"`
	}{
		Message: "Wrong ad format",
	})
	return
}

func (controller Ad) Delete(ctx *gin.Context) {
	stringId := ctx.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{
			Message: "bad ID",
		})
		return
	}

	err = controller.adService.DeleteAd(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{
			Message: "Something went wrong while deleting the ad number ID",
		})
		return
	}

	ctx.JSON(http.StatusNoContent, struct{}{})
}
