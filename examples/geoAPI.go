package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mirrr/go-sypexgeo"
	"time"
)

func main() {
	geo := sypexgeo.New("SxGeoCity.dat")
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Any("/spxgeo/:ip", func(c *gin.Context) {
		c.Header("Expires", time.Now().String())
		c.Header("Cache-Control", "no-cache")
		info, _ := geo.GetCityFull(c.Param("ip"))
		c.JSON(200, info)
	})

	fmt.Println("Server started at http://localhost:3000")
	router.Run(":3000")
}
