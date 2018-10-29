package ta

import (
	"errors"
	"strconv"
	"time"

	"github.com/oneleo/technical-analysis/calculate"
)

type RsiData struct {
	Time   time.Time
	Close  float64
	Result float64
	Signal int8
}

func (r *RsiData) ToString() []string {
	result := make([]string, 4)

	result[0] = r.Time.Format("2006-01-02T15:04:05Z07:00")
	//result[0] = strconv.FormatInt(r.Time.Unix(), 10)
	result[1] = strconv.FormatFloat(r.Close, 'E', -1, 64)
	result[2] = strconv.FormatFloat(r.Result, 'E', -1, 64)
	// 10 進制
	result[3] = strconv.FormatInt(int64(r.Signal), 10)

	return result
}

func Rsi(data []ICandle, baseDay int, upBound float64, downBound float64) (rsiData []RsiData, err error) {

	if len(data) < baseDay {
		return []RsiData{}, errors.New("The Data length is less than base day.")
	}
	if baseDay < 1 {
		return []RsiData{}, errors.New("The base day must be great than 1.")
	}
	if upBound <= downBound {
		return []RsiData{}, errors.New("The upBound is less than downBound.")
	}

	kLength := len(data)

	up := make([]float64, kLength)
	upAverage := make([]float64, kLength)
	down := make([]float64, kLength)
	downAverage := make([]float64, kLength)
	//signal = make([]int8, kLength)
	//result = make([]float64, kLength)

	rsidata := make([]RsiData, kLength)

	for i := 0; i < kLength; i++ {
		if i > 0 {
			if data[i].Array()[4] > data[i-1].Array()[4] {
				// 向上的價格變動
				up[i] = data[i].Array()[4] - data[i-1].Array()[4]
				down[i] = 0.0
			} else {
				up[i] = 0.0
				// 向下的價格變動
				down[i] = data[i-1].Array()[4] - data[i].Array()[4]
			}
		} else {
			up[i] = 0.0
			down[i] = 0.0
		}

		if i > baseDay-1 {
			upAverage[i] = calculate.Average(up[i-baseDay+1 : i+1])
			downAverage[i] = calculate.Average(down[i-baseDay+1 : i+1])
			// 計算 RSI 指標
			rsidata[i].Result = 100 - (100 / ((upAverage[i] / downAverage[i]) + 1.0))
			// 查看是否有碰觸邊界
			if rsidata[i].Result > upBound {
				rsidata[i].Signal = 1
			} else if rsidata[i].Result < downBound {
				rsidata[i].Signal = -1
			} else {
				rsidata[i].Signal = 0
			}

		} else {
			upAverage[i] = 0.0
			downAverage[i] = 0.0
			rsidata[i].Result = 0.0
			rsidata[i].Signal = 0
		}
		//result = result + data[i].Array()[4]
		rsidata[i].Time = time.Unix(int64(data[i].Array()[0]), int64(0))
		rsidata[i].Close = data[i].Array()[4]

	}

	return rsidata, nil
}
