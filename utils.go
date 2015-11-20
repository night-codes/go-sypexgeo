package sypexgeo

import (
	"encoding/binary"
	"errors"
	"math"
	"strconv"
	"strings"
)

func readUint32(bs []byte, offset int) uint32 { // faster than `binary.BigEndian.Uint32`
	return uint32(bs[offset])<<24 + uint32(bs[offset+1])<<16 + uint32(bs[offset+2])<<8 + uint32(bs[offset+3])
}
func readUint16(bs []byte, offset int) uint16 {
	return uint16(bs[offset])<<8 + uint16(bs[offset+1])
}
func readUint32L(bs []byte, offset int) uint32 {
	return uint32(bs[offset+3])<<24 + uint32(bs[offset+2])<<16 + uint32(bs[offset+1])<<8 + uint32(bs[offset])
}
func readUint16L(bs []byte, offset int) uint16 {
	return uint16(bs[offset+1])<<8 + uint16(bs[offset])
}
func readN32L(bs []byte, offset int, coma int) float32 {
	return float32(binary.LittleEndian.Uint32(bs[offset:offset+4])) / float32(math.Pow10(coma))
}
func readN16L(bs []byte, offset int, coma int) float32 {
	return float32(binary.LittleEndian.Uint16(bs[offset:offset+2])) / float32(math.Pow10(coma))
}
func readString(bs []byte, offset int) string { // faster than `binary.BigEndian.Uint32`
	i := offset
	for {
		i++
		if bs[i] == 0 || i >= len(bs) {
			break
		}
	}
	return string(bs[offset:i])
}

func fullUint32(bs []byte) (r []uint32) {
	r = make([]uint32, len(bs)/4)
	for i := 0; i < len(bs); i += 4 {
		r[i/4] = readUint32(bs, i)
	}
	return
}
func sliceUint32(bs []byte, offset, getLen uint32) uint32 {
	if getLen > 4 {
		panic("sliceUint32: getLen > 4")
	}
	return readUint32(append(make([]byte, 4-getLen), bs[offset:offset+getLen]...), 0)
}
func sliceUint32L(bs []byte, offset, getLen int) uint32 {
	if getLen > 4 {
		panic("sliceUint32: getLen > 4")
	}
	return readUint32L(append(append([]byte{}, bs[offset:offset+getLen]...), make([]byte, 4-getLen)...), 0)
}

func getIPByte(IP string, n byte) (byte, error) {
	sIP := strings.Split(IP, ".")
	if len(sIP) < 4 {
		return 0, errors.New("IP out of range")
	}
	b, err := strconv.ParseUint(sIP[n], 10, 8)
	if err != nil {
		return 0, err
	}
	return byte(b), nil
}

// Convert string IP to int64
func ipToN(IP string) int64 {
	b0, _ := getIPByte(IP, 0)
	b1, _ := getIPByte(IP, 1)
	b2, _ := getIPByte(IP, 2)
	b3, _ := getIPByte(IP, 3)
	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)
	return sum
}

func max32(a uint32, b uint32) (c uint32) {
	c = a
	if b > a {
		c = b
	}
	return
}

func min32(a uint32, b uint32) (c uint32) {
	c = a
	if b < a {
		c = b
	}
	return
}
