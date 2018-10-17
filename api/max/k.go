package max

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	maxapi "github.com/maicoin/max-exchange-api-go"
	"github.com/maicoin/max-exchange-api-go/models"
	"github.com/oneleo/technical-analysis/convert"
	"github.com/oneleo/technical-analysis/file"
)

func GetK1MinCsv(market string, dstFilePath string) (err error) {
	client := maxapi.NewClient()
	defer client.Close()

	var timestamp int32

	fileExist, err := file.IsExist(dstFilePath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if fileExist == true {
		fromCsv, err := file.CsvToArray(dstFilePath)
		if err != nil {
			fmt.Println(err)
			return err
		}

		candle, err := convert.StringsToMaxCandle(fromCsv[len(fromCsv)-1])
		if err != nil {
			fmt.Println(err)
			return err
		}
		// 因為目前 Max 交易所只能取得分鐘級別的 timestamp，所以將 timestamp 取得不大於它的分鐘級 timestamp。
		timestamp = int32(candle.Time.Unix())
		timestamp = timestamp - timestamp%60
		// 從下一分鐘開始。
		timestamp = timestamp + 60
	} else {
		timestamp = 0
		title := [][]string{{"Time", "Open", "High", "Low", "Close", "Volume"}}

		if err := file.ArrayToCsv(dstFilePath, title); err != nil {
			log.Fatalln(err)
		}
	}

	// 取得目前 Max 交易所時間
	currentTime, err := client.Time(context.Background())
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Max 交易所最小 timestamp = 1519206420。
	if timestamp < int32(1519206420) {
		timestamp = int32(1519206420)
	} else if timestamp > int32(currentTime.Unix()) {
		timestamp = int32(currentTime.Unix())
	}

	fmt.Println("TimeStanp: from ", timestamp, "（", time.Unix(int64(timestamp), int64(0)), "）\n\tto ", int32(currentTime.Unix()), "（", currentTime, "）")

	for t := int32(timestamp); t <= int32(currentTime.Unix()); t = t + 60 {
		k, err := GetK1Min(client, market, t)
		if err != nil {
			fmt.Println(err)
			return err
		}
		out := convert.MaxCandleToStrings(k)

		data := [][]string{out}
		err = file.ArrayAppendCsv(dstFilePath, data)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

// GetK1Min method will get OHLC(k line) of a specific market.
// Reference: https://max.maicoin.com/documents/api_list
// https://mholt.github.io/json-to-go/
func GetK1Min(client maxapi.API, market string, timestamp int32) (k models.Candle, err error) {
	// client := max.NewClient()
	// defer client.Close()

	// K returns OHLC chart of specific market.
	//
	// Available `CallOption`:
	//     Timestamp(): the seconds elapsed since Unix epoch, set to return data after the timestamp only
	//     Time(): the time in Go format, set to return data after the time only
	//     Period(): time period of K line in minute, default to 1
	//     PeriodDuration(): time period of K line in time.Duration format, default to 1*time.Minute
	//     Limit(): returned data points limit, default to 30
	// Reference:
	// https://github.com/maicoin/max-exchange-api-go/blob/master/interface.go
	results, err := client.K(context.Background(), market, maxapi.CallOption(func(opt map[string]interface{}) {
		opt["timestamp"] = timestamp
		opt["limit"] = int32(1)
		opt["period"] = int32(1)
	}))
	/*
		results, err := client.K(context.Background(), market, max.Timestamp(timestamp), max.Period(1), max.Limit(1))
	*/
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return models.Candle{}, err
	}

	resultBytes, err := json.MarshalIndent(results, "", "\t")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return models.Candle{}, err
	}

	var tmp []models.Candle
	json.Unmarshal(resultBytes, &tmp)
	k = tmp[0]

	return k, err
}
