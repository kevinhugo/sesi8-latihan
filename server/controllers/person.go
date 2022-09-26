package controllers

import (
	"net/http"
	"sesi8-latihan/server/models"
	"sesi8-latihan/server/views"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PersonController struct {
	db *gorm.DB
}

func NewPersonController(db *gorm.DB) *PersonController {
	return &PersonController{
		db: db,
	}
}

// GetPeople godoc
// @Summary Get all people
// @Decription Get list people
// @Tags person
// @Accept json
// @Produce json
// @Success 200 {object} views.GetAllPeopleSwagger
// @Router /people [get]
func (p *PersonController) GetPeople(ctx *gin.Context) {
	var people []models.Person

	err := p.db.Find(&people).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_PEOPLE_SUCCESS",
		Payload: people,
	})

}

// DeletePerson godoc
// @Summary Delete a person by id
// @Decription Delete a person by id
// @Tags person
// @Accept json
// @Produce json
// @Success 200 {object} views.GetAllPeopleSwagger
// @Router /people [delete]
func (p *PersonController) DeletePerson(ctx *gin.Context) {
	var people []models.Person

	err := p.db.Find(&people).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_PEOPLE_SUCCESS",
		Payload: people,
	})

}

// AddPerson godoc
// @Summary Add a person
// @Decription Add a person
// @Tags person
// @Accept json
// @Produce json
// @Success 200 {object} views.GetAllPeopleSwagger
// @Router /people [post]
func (p *PersonController) AddPerson(ctx *gin.Context) {
	var people []models.Person

	err := p.db.Find(&people).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_PEOPLE_SUCCESS",
		Payload: people,
	})

}

// UpdatePerson godoc
// @Summary Update person
// @Decription Update person
// @Tags person
// @Accept json
// @Produce json
// @Success 200 {object} views.GetAllPeopleSwagger
// @Router /people [put]
func (p *PersonController) UpdatePerson(ctx *gin.Context) {
	var people []models.Person

	err := p.db.Find(&people).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_PEOPLE_SUCCESS",
		Payload: people,
	})

}
