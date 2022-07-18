package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"base-app/models"
	"base-app/services"
)
type PetController struct {
	PetService services.PetService
}

func New(petservice services.PetService) PetController {
	return PetController{
		PetService: petservice,
	}
}

func (pc *PetController) CreatePet(ctx *gin.Context) {
	var pet models.Pet
	if err := ctx.ShouldBindJSON(&pet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Invalid Input": err.Error()})
		return
	}
	err := pc.PetService.CreatePet(&pet)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (pc *PetController) GetPet(ctx *gin.Context) {
	var id string = ctx.Param("id")
	pet, err := pc.PetService.GetPet(&id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,pet)
}

func (pc *PetController) GetAll(ctx *gin.Context) {
	pets, err := pc.PetService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,pets)
}

func (pc *PetController) UpdatePet(ctx *gin.Context) {
	var pet models.Pet
	if err := ctx.ShouldBindJSON(&pet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := pc.PetService.UpdatePet(&pet)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (pc *PetController) DeletePet(ctx *gin.Context) {
	var id string = ctx.Param("id")
	err := pc.PetService.DeletePet(&id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}


	

