package ta

import (
	"strconv"
)

type Price = float64

type Volume = float64

type Timestamp = int32

func ParsePrice(s string) (Price, error) {
	return strconv.ParseFloat(s, 64)
}

func ParseVolume(s string) (Volume, error) {
	return strconv.ParseFloat(s, 64)
}
