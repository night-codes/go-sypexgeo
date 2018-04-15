package main

import (
	"encoding/json"
	"fmt"
	"github.com/night-codes/go-sypexgeo"
)

func main() {
	geo := sypexgeo.New("../SxGeoCity.dat")
	printJSON(geo.Info("93.73.35.74"))
}

func printJSON(t ...interface{}) {
	j, _ := json.MarshalIndent(t, "", "\t")
	fmt.Println(string(j))
}
