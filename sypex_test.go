package sypexgeo

import (
	"fmt"
	"testing"
)

var geo SxGEO

func init() {
	geo = New("examples/SxGeoCity.dat")
}

func TestErrors(t *testing.T) {
	fmt.Print("TestErrors... ")
	_, err := geo.GetCityFull("0.0.0.0")
	if err == nil {
		t.Error("Expected error object when used incorrect arguments")
	}
	fmt.Println("PASS")
}

func TestGetCityFull1(t *testing.T) {
	fmt.Print("TestGetCityFull(1)... ")
	info, err := geo.GetCityFull("93.73.35.74")
	if err != nil {
		t.Error(err)
	}
	if info["city"].(Obj)["name_en"].(string) != "Kiev" {
		t.Error("`name_en` not equal `Kiev`:", info["city"].(Obj)["name_en"])
	}
	if info["region"].(Obj)["name_en"].(string) != "Kyiv" {
		t.Error("`name_en` not equal `Kyiv`", info["region"].(Obj)["name_en"])
	}
	_, err = geo.GetCityFull("0.0.0.0")
	if err == nil {
		t.Error("Expected error object when used incorrect arguments")
	}
	fmt.Println("PASS")
}

func TestGetCityFull2(t *testing.T) {
	fmt.Print("TestGetCityFull(2)... ")
	info, err := geo.GetCityFull("124.13.35.14")
	if err != nil {
		t.Error(err)
	}
	if info["country"].(Obj)["name_en"].(string) != "Malaysia" {
		t.Error("`name_en` not equal `Malaysia`:", info["country"].(Obj)["name_en"])
	}
	fmt.Println("PASS")
}

func TestGetCity(t *testing.T) {
	fmt.Print("TestGetCity... ")
	info, err := geo.GetCity("17.21.34.42")
	if err != nil {
		t.Error(err)
	}
	if info["city"].(Obj)["name_en"].(string) != "Cupertino" {
		t.Error("`name_en` not equal `Cupertino`", info["city"].(Obj)["name_en"])
	}
	fmt.Println("PASS")
}

func TestGetCountry(t *testing.T) {
	fmt.Print("TestGetCountry... ")
	code, err := geo.GetCountry("78.120.24.150")
	if err != nil {
		t.Error(err)
	}
	if code != "FR" {
		t.Error("Country code not equal `FR`", code)
	}
	fmt.Println("PASS")
}

func TestGetCountryID(t *testing.T) {
	fmt.Print("TestGetCountryID... ")
	id, err := geo.GetCountryID("111.12.62.12")
	if err != nil {
		t.Error(err)
	}
	if id != 48 {
		t.Error("Country id not equal 48", id)
	}
	fmt.Println("PASS")
}
