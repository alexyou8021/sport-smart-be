package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/alexyou8021/sport-smart-be/controllers"
)

func GetEvents(ctx *gin.Context) {
	result := controllers.GetEvents()
	ctx.JSON(http.StatusOK, result)
}

func GetMembers(ctx *gin.Context) {
	result := controllers.GetMembers()
	ctx.JSON(http.StatusOK, result)
}
