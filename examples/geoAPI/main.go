package main

import (
	"net/http"
	"time"

	"github.com/night-codes/go-sypexgeo"
	"github.com/night-codes/tokay"
)

func main() {
	geo := sypexgeo.New("../SxGeoCity.dat")
	router := tokay.New()

	router.Any("/spxgeo/<ip>", func(c *tokay.Context) {
		c.Header("Expires", time.Now().String())
		c.Header("Cache-Control", "no-cache")
		info, err := geo.Info(c.Param("ip"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		c.JSON(200, info)
	})

	router.Run(":3000", "Server started. Try: http://localhost%s/spxgeo/1.4.10.10")
}
