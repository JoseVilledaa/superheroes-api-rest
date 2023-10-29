package controllers

import (
	"net/http"

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

func (sc *SuperheroeController) RegisterRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/superheroes")
	route.GET("/getall", sc.GetAll)
}
