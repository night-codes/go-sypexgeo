package main

import (
	"encoding/json"
	"fmt"
	"github.com/mirrr/sypexgeo"
)

func main() {
	geo := sypexgeo.New("SxGeoCity.dat")
	printJSON(geo.GetCityFull("93.73.35.74"))
}

func printJSON(t ...interface{}) {
	j, _ := json.MarshalIndent(t, "", "\t")
	fmt.Println(string(j))
}
