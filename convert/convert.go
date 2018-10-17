package convert

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/maicoin/max-exchange-api-go/models"
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

func MaxCandleToStrings(k models.Candle) (out []string) {
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

func StringsToMaxCandle(in []string) (k models.Candle, err error) {

	if len(in) != 6 {
		return models.Candle{}, errors.New("The length of the input is incorrect.")
	}
	t, err := StringToInt64(in[0])
	if err != nil {
		return models.Candle{}, errors.New(fmt.Sprint("The \"Time\" of the input is incorrect. - ", err.Error()))
	}
	open, err := StringToFloat64(in[1])
	if err != nil {
		return models.Candle{}, errors.New(fmt.Sprint("The \"Open\" of the input is incorrect. - ", err.Error()))
	}

	high, err := StringToFloat64(in[2])
	if err != nil {
		return models.Candle{}, errors.New(fmt.Sprint("The \"High\" of the input is incorrect. - ", err.Error()))
	}

	low, err := StringToFloat64(in[3])
	if err != nil {
		return models.Candle{}, errors.New(fmt.Sprint("The \"Low\" of the input is incorrect. - ", err.Error()))
	}

	close, err := StringToFloat64(in[4])
	if err != nil {
		return models.Candle{}, errors.New(fmt.Sprint("The \"Close\" of the input is incorrect. - ", err.Error()))
	}

	volume, err := StringToFloat64(in[5])
	if err != nil {
		return models.Candle{}, errors.New(fmt.Sprint("The \"Volume\" of the input is incorrect. - ", err.Error()))
	}

	return models.Candle{time.Unix(t, int64(0)), open, high, low, close, volume}, nil
}
