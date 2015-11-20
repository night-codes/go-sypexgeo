// Package sypexgeo to access data from Sypex Geo IP database files,
// accepts only SypexGeo 2.2 databases
package sypexgeo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type (
	// Obj is wrap for map[string]interface{}
	Obj      map[string]interface{}
	dbSlices struct {
		BIndex    []uint32 // индекс первого байта
		MIndex    []uint32 // главный индекс
		DB        []byte
		Regions   []byte
		Cities    []byte
		Countries []byte
	}
	finder struct {
		BLen       uint32 // кол-во элементов индекса первого байта
		MLen       uint32 // кол-во элементов главного индекса
		Range      uint32 // Блоков в одном элементе индекса (до 65 тыс.)
		DBItems    uint32 // Количество диапазонов в базе (айпишников)
		IDLen      uint32 // Размер ID-блока 1-для городов, 3-для стран
		BlockLen   uint32 // Размер блока BD = IDLen+3
		PackLen    uint32 // Размер блока описания упаковки
		MaxRegion  uint32 // максимальный размер записи в справочнике регионов
		MaxCity    uint32 // максимальный размер записи в справочнике городов
		MaxCountry uint32 // максимальный размер записи в справочнике стран
		CountryLen uint32 // размер справочника стран
		Pack       []string
		S          dbSlices
	}
	// SxGEO main object
	SxGEO struct {
		finder  finder
		Version float32
		Updated uint32
	}
)

func (f *finder) getLocationOffset(IP string) (uint32, error) {
	firstByte, err := getIPByte(IP, 0)
	if err != nil {
		return 0, err
	}
	IPn := uint32(ipToN(IP))
	if firstByte == 0 || firstByte == 10 || firstByte == 127 || uint32(firstByte) >= f.BLen || IPn == 0 {
		return 0, errors.New("IP out of range")
	}

	var min, max uint32
	minIndex, maxIndex := f.S.BIndex[firstByte-1], f.S.BIndex[firstByte]

	if maxIndex-minIndex > f.Range {
		part := f.searchIdx(IPn, minIndex/f.Range, maxIndex/f.Range-1)
		max = f.DBItems
		if part > 0 {
			min = part * f.Range
		}
		if part <= f.MLen {
			max = (part + 1) * f.Range
		}
		min, max = max32(min, minIndex), min32(max, maxIndex)
	} else {
		min, max = minIndex, maxIndex
	}
	return f.searchDb(IPn, min, max), nil
}

func (f *finder) searchDb(IPn, min, max uint32) uint32 {
	if max-min > 1 {
		IPn &= 0x00FFFFFF

		for max-min > 8 {
			offset := (min + max) >> 1
			// if IPn > substr(str, offset*f.block_len, 3) {
			if IPn > sliceUint32(f.S.DB, offset*f.BlockLen, 3) {
				min = offset
			} else {
				max = offset
			}
		}

		for IPn >= sliceUint32(f.S.DB, min*f.BlockLen, 3) {
			min++
			if min >= max {
				break
			}
		}
	} else {
		min++
	}

	return sliceUint32(f.S.DB, min*f.BlockLen-f.IDLen, f.IDLen)
}

func (f *finder) searchIdx(IPn, min, max uint32) uint32 {
	var offset uint32
	if max < min {
		max, min = min, max
	}
	for max-min > 8 {
		offset = (min + max) >> 1
		if IPn > uint32(f.S.MIndex[offset]) {
			min = offset
		} else {
			max = offset
		}
	}
	for IPn > uint32(f.S.MIndex[min]) && min <= max {
		min++
	}
	return min
}

func (f *finder) unpack(seek, uType uint32) (Obj, error) {
	var bs []byte
	var maxLen uint32
	ret := Obj{}

	if int(uType+1) > len(f.Pack) {
		return Obj{}, errors.New("Pack method not found")
	}

	switch uType {
	case 0:
		bs = f.S.Cities
		maxLen = f.MaxCountry
	case 1:
		bs = f.S.Regions
		maxLen = f.MaxRegion
	case 2:
		bs = f.S.Cities
		maxLen = f.MaxCity
	}

	raw := bs[seek : seek+maxLen]

	var cursor int
	for _, el := range strings.Split(f.Pack[uType], "/") {
		cmd := strings.Split(el, ":")
		switch string(cmd[0][0]) {
		case "T":
			ret[cmd[1]] = raw[cursor]
			cursor++
		case "M":
			ret[cmd[1]] = sliceUint32L(raw, cursor, 3)
			cursor += 3
		case "S":
			ret[cmd[1]] = readUint16L(raw, cursor)
			cursor += 2
		case "b":
			ret[cmd[1]] = readString(raw, cursor)
			cursor += len(ret[cmd[1]].(string)) + 1
		case "c":
			if len(cmd[0]) > 1 {
				ln, _ := strconv.Atoi(string(cmd[0][1:]))
				ret[cmd[1]] = string(raw[cursor : cursor+ln])
				cursor += ln
			}
		case "N":
			if len(cmd[0]) > 1 {
				coma, _ := strconv.Atoi(string(cmd[0][1:]))
				ret[cmd[1]] = readN32L(raw, cursor, coma)
				cursor += 4
			}
		case "n":
			if len(cmd[0]) > 1 {
				coma, _ := strconv.Atoi(string(cmd[0][1:]))
				ret[cmd[1]] = readN16L(raw, cursor, coma)
				cursor += 2
			}
		}
	}
	return ret, nil
}

func (f *finder) parseCity(seek uint32, full bool) (Obj, error) {
	if f.PackLen == 0 {
		return Obj{}, errors.New("Pack methods not found")
	}
	country, city, region := Obj{}, Obj{}, Obj{}
	var err error
	onlyCountry := false

	if seek < f.CountryLen {
		country, err = f.unpack(seek, 0)
		city = Obj{
			"id":      0,
			"lat":     country["lat"],
			"lon":     country["lon"],
			"name_en": "",
			"name_ru": "",
		}
		onlyCountry = true
	} else {
		city, err = f.unpack(seek, 2)
		country = Obj{"id": city["country_id"], "iso": isoCodes[city["country_id"].(uint8)]}
		delete(city, "country_id")
	}

	if err != nil {
		return Obj{}, err
	}

	if full {
		_ = onlyCountry
		if !onlyCountry {
			region, err = f.unpack(city["region_seek"].(uint32), 1)
			if err != nil {
				return Obj{}, err
			}
			country, err = f.unpack(uint32(region["country_seek"].(uint16)), 0)
			delete(city, "region_seek")
			delete(region, "country_seek")
		}

		return Obj{"country": country, "region": region, "city": city}, err
	}

	delete(city, "region_seek")
	return Obj{"country": country, "region": region, "city": city}, err
}

// GetCountry return string country iso-code, like `RU`, `UA` etc.
func (s *SxGEO) GetCountry(IP string) (string, error) {
	info, err := s.GetCity(IP)
	if err != nil {
		return "", err
	}
	fmt.Println(info["country"].(Obj))

	return info["country"].(Obj)["iso"].(string), nil
}

// GetCountryID return integer country identifier
func (s *SxGEO) GetCountryID(IP string) (int, error) {
	info, err := s.GetCity(IP)
	if err != nil {
		return 0, err
	}
	fmt.Println(info["country"].(Obj))

	return int(info["country"].(Obj)["id"].(uint8)), nil
}

// GetCityFull get full info by IP (with regions and contries data)
func (s *SxGEO) GetCityFull(IP string) (map[string]interface{}, error) {
	seek, err := s.finder.getLocationOffset(IP)
	if err != nil {
		return Obj{}, err
	}
	return s.finder.parseCity(seek, true)
}

// GetCity get short info by IP
func (s *SxGEO) GetCity(IP string) (map[string]interface{}, error) {
	seek, err := s.finder.getLocationOffset(IP)
	if err != nil {
		return Obj{}, err
	}
	return s.finder.parseCity(seek, false)
}

// New SxGEO object
func New(filename string) SxGEO {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Database file not found")
	} else if string(dat[:3]) != `SxG` && dat[3] != 22 && dat[8] != 2 {
		panic("Wrong database format")
	} else if dat[9] != 0 {
		panic("Only UTF-8 version is supported")
	}

	IDLen := uint32(dat[19])
	blockLen := 3 + IDLen                   // размер блока в базе диапазонов
	DBItems := readUint32(dat, 15)          // кол-во диапазонов IP
	BLen := uint32(dat[10])                 // количество элементов в индексе первого байта
	MLen := uint32(readUint16(dat, 11))     // количество элементов в главном индексе
	packLen := uint32(readUint16(dat, 38))  // Размер описания формата упаковки города/региона/страны
	regnLen := readUint32(dat, 24)          // Размер справочника регионов
	cityLen := readUint32(dat, 28)          // Размер справочника городов
	countryLen := readUint32(dat, 34)       // Размер справочника стран
	BStart := uint32(40 + packLen)          // Начало индекса первого байта
	MStart := BStart + BLen*4               // Начало главного индекса
	DBStart := MStart + MLen*4              // Начало базы диапазонов
	regnStart := DBStart + DBItems*blockLen // Начало справочника регионов
	cityStart := regnStart + regnLen        // Начало справочника городов
	cntrStart := cityStart + cityLen        // Начало справочника стран
	cntrEnd := cntrStart + countryLen
	pack := strings.Split(string(dat[40:40+packLen]), string(byte(0)))

	return SxGEO{
		Version: float32(dat[3]) / 10,
		Updated: readUint32(dat, 4),
		finder: finder{
			Range:      uint32(readUint16(dat, 13)),
			DBItems:    DBItems,
			IDLen:      IDLen,
			BLen:       BLen,
			MLen:       MLen,
			CountryLen: countryLen,
			BlockLen:   blockLen,
			PackLen:    packLen,
			Pack:       pack,
			MaxRegion:  uint32(readUint16(dat, 20)),
			MaxCity:    uint32(readUint16(dat, 22)),
			MaxCountry: uint32(readUint16(dat, 32)),
			S: dbSlices{
				BIndex:    fullUint32(dat[BStart:MStart]),
				MIndex:    fullUint32(dat[MStart:DBStart]),
				DB:        dat[DBStart:regnStart],
				Regions:   dat[regnStart:cityStart],
				Cities:    dat[cityStart:cntrStart],
				Countries: dat[cntrStart:cntrEnd],
			},
		},
	}
}
