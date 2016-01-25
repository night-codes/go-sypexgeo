package main

import (
	"fmt"
	"gopkg.in/mirrr/go-sypexgeo.v1"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	geo := sypexgeo.New("SxGeoCity.dat")

	t := time.Now()
	for i := 0; i < 100000; i++ {
		ip := strconv.Itoa(rand.Intn(222)) + "." + strconv.Itoa(rand.Intn(222)) + "." + strconv.Itoa(rand.Intn(222)) + "." + strconv.Itoa(rand.Intn(222))
		geo.GetCityFull(ip)
	}
	fmt.Printf("The call took %v to run.\n", time.Now().Sub(t))
}
