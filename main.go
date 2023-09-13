package main

import (
	"github/meshachdamilare/hng_stage_two/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := routers.NewRouter()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "server works")
	})

	router.Run(":3000")
}
