package main

import (
	"encoding/json"
	"fmt"
	"github.com/night-codes/go-sypexgeo"
	"github.com/night-codes/tokay"
	"time"
)

func main() {
	geo := sypexgeo.New("../SxGeoCity.dat")
	router := tokay.New()

	router.Any("/spxgeo/<ip>", func(c *tokay.Context) {
		c.Header("Expires", time.Now().String())
		c.Header("Cache-Control", "no-cache")
		info, _ := geo.Info(c.Param("ip"))
		j, _ := json.MarshalIndent(info, "", "\t")
		c.String(200, string(j))
	})

	fmt.Println("Server started at http://localhost:3000")
	router.Run(":3000")
}
