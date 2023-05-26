package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/alexyou8021/sport-smart-be/controllers"
)

func GetEvents(ctx *gin.Context) {
	result := controllers.Controller()
	ctx.String(http.StatusOK, result)
}
