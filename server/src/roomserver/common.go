package main

import (
	"math/rand"
	"strconv"
	"strings"
)

func StrToInt64(s string) int64 {
	if s == "" {
		return 0
	}
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i64
}

func StrToUint64(s string) uint64 {
	if s == "" {
		return 0
	}
	ui64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return ui64
}

func StrToInt32(s string) int32 {
	if s == "" {
		return 0
	}
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return int32(i64)
}

func StrToInt(s string) int {
	if s == "" {
		return 0
	}
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return int(i64)
}

func StrToUint16(s string) uint16 {
	if s == "" {
		return 0
	}
	i64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return uint16(i64)
}

func StrToUint32(s string) uint32 {
	if s == "" {
		return 0
	}
	i64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return uint32(i64)
}

func StrToFloat32(s string) float32 {
	if s == "" {
		return 0
	}
	f64, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0
	}
	return float32(f64)
}

func StrToFloat64(s string) float64 {
	if s == "" {
		return 0
	}
	f64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f64
}

func StrToUint32List(s, sp string) []uint32 {
	list := make([]uint32, 0)
	if s == "" {
		return list
	}
	for _, v := range strings.Split(s, sp) {
		i64, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			//return make([]uint32, 0)
			continue
		}
		list = append(list, uint32(i64))
	}
	return list
}

func RandToken(n int) string {
	const allowed = "abcdefghijklmnopqrstuvwxyz1234567890"
	b := make([]byte, n)
	for i := range b {
		b[i] = allowed[rand.Intn(len(allowed))]
	}

	return string(b)
}
