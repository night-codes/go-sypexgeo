package main

import (
	"fmt"
	"github.com/night-codes/go-sypexgeo"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	geo := sypexgeo.New("../SxGeoCity.dat")

	t := time.Now()
	for i := 0; i < 1000000; i++ {
		ip := strconv.Itoa(rand.Intn(222)) + "." + strconv.Itoa(rand.Intn(222)) + "." + strconv.Itoa(rand.Intn(222)) + "." + strconv.Itoa(rand.Intn(222))
		geo.Info(ip)
	}
	fmt.Printf("The call took %v to run.\n", time.Now().Sub(t))
}
