package models

import (
	"time"

	"github.com/oneleo/technical-analysis/api/coinmarketcap/types"
)

type Candle struct {
	Time   time.Time   `json:"timestamp,omitempty"`
	Open   types.Price `json:"open,omitempty"`
	High   types.Price `json:"high,omitempty"`
	Low    types.Price `json:"low,omitempty"`
	Close  types.Price `json:"close,omitempty"`
	Volume types.Price `json:"volume,omitempty"`
	Cap    types.Price `json:"cap,omitempty"`
}

func (c *Candle) Array() []float64 {
	result := make([]float64, 7)
	result[0] = float64(c.Time.Unix())
	result[1] = c.Open
	result[2] = c.High
	result[3] = c.Low
	result[4] = c.Close
	result[5] = c.Volume
	result[6] = c.Cap

	return result
}
