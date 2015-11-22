package main

/*
typedef struct {
	int id;
	float lat;
	float lon;
	char* name_en;
	char* name_ru;
} spxCity;

typedef struct {
	int id;
	char* iso;
	char* name_en;
	char* name_ru;
} spxRegion;

typedef struct {
	int id;
	char* iso;
	float lat;
	float lon;
	char* name_en;
	char* name_ru;
} spxCountry;

typedef struct {
	spxCity city;
	spxCountry country;
	spxRegion region;
	int err;
} spxInfo;
*/
import "C"
import (
	"github.com/mirrr/go-sypexgeo"
	"github.com/mirrr/types"
)

var geo sypexgeo.SxGEO

//export spxConnect
func spxConnect(c *C.char) {
	filepath := C.GoString(c)
	geo = sypexgeo.New(filepath)
}

//export spxGetCity
func spxGetCity(ip *C.char, full C.int) (ret C.spxInfo) {
	var info sypexgeo.Obj
	var err error
	if int(full) == 1 {
		info, err = geo.GetCityFull(C.GoString(ip))
	} else {
		info, err = geo.GetCity(C.GoString(ip))
	}

	if err != nil {
		ret.err = 1
	} else {
		city, _ := info["city"].(sypexgeo.Obj)
		country, _ := info["country"].(sypexgeo.Obj)
		region, _ := info["region"].(sypexgeo.Obj)

		ret = C.spxInfo{
			city: C.spxCity{
				id:      C.int(types.Int(city["id"])),
				lat:     C.float(types.Float32(city["lat"])),
				lon:     C.float(types.Float32(city["lon"])),
				name_en: C.CString(types.String(city["name_en"])),
				name_ru: C.CString(types.String(city["name_ru"])),
			},
			country: C.spxCountry{
				id:      C.int(types.Int(country["id"])),
				iso:     C.CString(types.String(country["iso"])),
				lat:     C.float(types.Float32(country["lat"])),
				lon:     C.float(types.Float32(country["lon"])),
				name_en: C.CString(types.String(country["name_en"])),
				name_ru: C.CString(types.String(country["name_ru"])),
			},
			region: C.spxRegion{
				id:      C.int(types.Int(region["id"])),
				iso:     C.CString(types.String(region["iso"])),
				name_en: C.CString(types.String(region["name_en"])),
				name_ru: C.CString(types.String(region["name_ru"])),
			},
		}
	}
	return
}

//export spxGetCountry
func spxGetCountry(ip *C.char) *C.char {
	iso, _ := geo.GetCountry(C.GoString(ip))
	return C.CString(iso)
}

//export spxGetCountryID
func spxGetCountryID(ip *C.char) C.int {
	id, _ := geo.GetCountryID(C.GoString(ip))
	return C.int(id)
}

func main() {
}
