package ta

import (
	"strconv"
	"time"

	"github.com/oneleo/technical-analysis/api/coinmarketcap/types"
)

type Candle struct {
	Time   time.Time
	Open   types.Price
	High   types.Price
	Low    types.Price
	Close  types.Price
	Volume types.Price
	Cap    types.Price
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

func (c *Candle) ToString() []string {
	result := make([]string, 7)

	result[0] = strconv.FormatInt(c.Time.Unix(), 10)
	//result[0] = c.Time.Format("2006-01-02T15:04:05Z07:00")
	result[1] = strconv.FormatFloat(c.Open, 'E', -1, 64)
	result[2] = strconv.FormatFloat(c.High, 'E', -1, 64)
	result[3] = strconv.FormatFloat(c.Low, 'E', -1, 64)
	result[4] = strconv.FormatFloat(c.Close, 'E', -1, 64)
	result[5] = strconv.FormatFloat(c.Volume, 'E', -1, 64)
	result[6] = strconv.FormatFloat(c.Cap, 'E', -1, 64)

	return result
}
