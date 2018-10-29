package convert

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	maxModel "github.com/maicoin/max-exchange-api-go/models"
	"github.com/oneleo/technical-analysis/ta"
)

//type MaxCandle models.Candle
//type MaxCandles []MaxCandle

func StructInfo(data interface{}) string {
	v := reflect.ValueOf(data)
	n := v.NumField()

	st := reflect.TypeOf(data)
	headers := make([]string, n)
	for i := 0; i < n; i++ {
		headers[i] = fmt.Sprintf(`"%s": %d`, st.Field(i).Name, i)
	}

	rowContents := make([]string, n)
	for i := 0; i < n; i++ {
		x := v.Field(i)
		s := fmt.Sprintf("%v", x.Interface())
		if x.Type().String() == "string" {
			s = `"` + s + `"`
		}
		rowContents[i] = s
	}

	return "{" + strings.Join(headers, ", ") + `, "rows": [[` + strings.Join(rowContents, ", ") + "]]}"
}

func MaxCandleToStrings(k maxModel.Candle) (out []string) {
	out = append(out, strconv.FormatInt(k.Time.Unix(), 10))
	out = append(out, strconv.FormatFloat(k.Open, 'E', -1, 64))
	out = append(out, strconv.FormatFloat(k.High, 'E', -1, 64))
	out = append(out, strconv.FormatFloat(k.Low, 'E', -1, 64))
	out = append(out, strconv.FormatFloat(k.Close, 'E', -1, 64))
	out = append(out, strconv.FormatFloat(k.Volume, 'E', -1, 64))
	return out
}

func StringToFloat64(in string) (out float64, err error) {
	return strconv.ParseFloat(in, 64)
}

func StringToInt64(in string) (out int64, err error) {
	return strconv.ParseInt(in, 10, 64)
}

// TimeStringToInt64，要注意可能會因為時差，而導致轉出來的天數是隔天，但是是正常的。
func TimeStringToInt64(in string) (out int64, err error) {
	in = in + " 23:59:59 +0800 UTC"
	// Parse 請按照 01 月、02 日、15（03）時、04 分、05 秒、2006（06）年來分類。
	t, err := time.Parse("02-01-2006 15:04:05 -0700 UTC", in)
	return t.Unix(), err
}

func StringsToMaxCandle(in []string) (k maxModel.Candle, err error) {

	if len(in) != 6 {
		return maxModel.Candle{}, errors.New("The length of the input is incorrect.")
	}
	t, err := StringToInt64(in[0])
	if err != nil {
		return maxModel.Candle{}, errors.New(fmt.Sprint("The \"Time\" of the input is incorrect. - ", err.Error()))
	}
	open, err := StringToFloat64(in[1])
	if err != nil {
		return maxModel.Candle{}, errors.New(fmt.Sprint("The \"Open\" of the input is incorrect. - ", err.Error()))
	}

	high, err := StringToFloat64(in[2])
	if err != nil {
		return maxModel.Candle{}, errors.New(fmt.Sprint("The \"High\" of the input is incorrect. - ", err.Error()))
	}

	low, err := StringToFloat64(in[3])
	if err != nil {
		return maxModel.Candle{}, errors.New(fmt.Sprint("The \"Low\" of the input is incorrect. - ", err.Error()))
	}

	close, err := StringToFloat64(in[4])
	if err != nil {
		return maxModel.Candle{}, errors.New(fmt.Sprint("The \"Close\" of the input is incorrect. - ", err.Error()))
	}

	volume, err := StringToFloat64(in[5])
	if err != nil {
		return maxModel.Candle{}, errors.New(fmt.Sprint("The \"Volume\" of the input is incorrect. - ", err.Error()))
	}

	return maxModel.Candle{time.Unix(t, int64(0)), open, high, low, close, volume}, nil
}

// StringsToTaCandle, Standard TA Candle = {Time, Open, High, Low, Close, Volume, Cap}
func StringsToTaCandle(in []string) (k *ta.Candle, err error) {
	if len(in) != 7 {
		return &ta.Candle{}, errors.New("The length of the input is incorrect.")
	}

	var t int64
	var open, high, low, close, volume, cap float64

	if strings.Compare(in[0], "") == 0 {
		t = 0
	} else {
		t, err = TimeStringToInt64(in[0])
		if err != nil {
			return &ta.Candle{}, errors.New(fmt.Sprint("The \"Time\" of the input is incorrect. - ", err.Error()))
		}
	}

	if strings.Compare(in[1], "") == 0 {
		open = 0.0
	} else {
		open, err = StringToFloat64(in[1])
		if err != nil {
			return &ta.Candle{}, errors.New(fmt.Sprint("The \"Open\" of the input is incorrect. - ", err.Error()))
		}
	}

	if strings.Compare(in[2], "") == 0 {
		high = 0.0
	} else {
		high, err = StringToFloat64(in[2])
		if err != nil {
			return &ta.Candle{}, errors.New(fmt.Sprint("The \"High\" of the input is incorrect. - ", err.Error()))
		}
	}

	if strings.Compare(in[3], "") == 0 {
		low = 0.0
	} else {
		low, err = StringToFloat64(in[3])
		if err != nil {
			return &ta.Candle{}, errors.New(fmt.Sprint("The \"Low\" of the input is incorrect. - ", err.Error()))
		}
	}

	if strings.Compare(in[4], "") == 0 {
		close = 0.0
	} else {
		close, err = StringToFloat64(in[4])
		if err != nil {
			return &ta.Candle{}, errors.New(fmt.Sprint("The \"Close\" of the input is incorrect. - ", err.Error()))
		}
	}

	if strings.Compare(in[5], "") == 0 {
		volume = 0.0
	} else {
		volume, err = StringToFloat64(in[5])
		if err != nil {
			return &ta.Candle{}, errors.New(fmt.Sprint("The \"Volume\" of the input is incorrect. - ", err.Error()))
		}
	}

	if strings.Compare(in[6], "") == 0 {
		cap = 0.0
	} else {
		cap, err = StringToFloat64(in[6])
		if err != nil {
			return &ta.Candle{}, errors.New(fmt.Sprint("The \"Volume\" of the input is incorrect. - ", err.Error()))
		}
	}

	return &ta.Candle{time.Unix(t, int64(0)), open, high, low, close, volume, cap}, nil
}
