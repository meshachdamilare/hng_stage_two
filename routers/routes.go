package routers

import (
	"github/meshachdamilare/hng_stage_two/handlers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/api", handlers.CreatePerson)
	router.GET("/api/:id", handlers.GetPerson)
	router.PATCH("/api/:id", handlers.UpdatePerson)
	router.DELETE("/api/:id", handlers.DeletePerson)
	return router
}
