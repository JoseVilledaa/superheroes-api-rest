package controllers

import (
	"net/http"

	"github.com/JoseVilledaa/superheroes-api/models"
	"github.com/google/uuid"

	"github.com/JoseVilledaa/superheroes-api/services"
	"github.com/gin-gonic/gin"
)

type SuperheroeController struct {
	superheroeService services.SuperheroeService
}

func New(superheroeservice services.SuperheroeService) SuperheroeController {
	return SuperheroeController{
		superheroeService: superheroeservice,
	}

}

func (sc *SuperheroeController) GetAll(ctx *gin.Context) {
	superheroes, err := sc.superheroeService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, superheroes)
}

func (sc *SuperheroeController) CreateSuperheroe(ctx *gin.Context) {
	var superheroe models.Superheroe
	if err := ctx.ShouldBindJSON(&superheroe); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Genera un nuevo UUID
	id, err := uuid.NewRandom()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate UUID"})
		return
	}

	// Asigna el UUID al campo Id del Superhéroe
	superheroe.Id = id

	if err := sc.superheroeService.CreateSuperheroe(&superheroe); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Superheroe created successfully"})
}

func (sc *SuperheroeController) DeleteSuperheroe(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := sc.superheroeService.DeleteSuperheroe(id); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Superheroe deleted successfully"})
}

func (sc *SuperheroeController) RegisterRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/superheroes")
	route.GET("/getall", sc.GetAll)
	route.POST("/create", sc.CreateSuperheroe)
	route.DELETE("/delete/:id", sc.DeleteSuperheroe)
}
