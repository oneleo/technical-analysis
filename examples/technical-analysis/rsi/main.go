package main

import (
	"log"
	"strconv"

	_ "github.com/maicoin/max-exchange-api-go/models"
	"github.com/oneleo/technical-analysis/convert"
	"github.com/oneleo/technical-analysis/file"
	"github.com/oneleo/technical-analysis/ta"
)

func main() {

	coin := []string{"btc", "eth", "xrp", "bch", "eos",
		"xlm", "ltc", "usdt", "ada", "xmr", "trx", "miota", "dash", "bnb", "neo", "etc", "xem", "xtz", "zec", "vet", "doge", "btg", "mkr", "omg", "zrx", "dcr", "qtum", "ont", "lsk", "ae", "zil", "bcd", "nano", "sc", "bts", "icx", "bat", "dgb", "bcn", "xvg", "steem", "npxs", "waves", "btm", "tusd"}
	baseDay := []int{9, 14, 25}

	for _, c := range coin {
		for _, d := range baseDay {
			rsi(c, d)
		}
	}

}

func rsi(coin string, baseDay int) {
	fromCsv, err := file.CsvToArray("../../../history/coinmarketcap_" + coin + "_usd_k.csv")
	if err != nil {
		log.Fatalln(err)
	}

	// 去掉第一列的 Title
	length := int(len(fromCsv) - 1)
	fromCsvReverse := make([][]string, length)

	for i := 0; i < length; i++ {
		fromCsvReverse[length-1-i] = fromCsv[i+1]
	}

	k := make([]ta.ICandle, len(fromCsvReverse))

	for i := 0; i < len(k); i++ {
		k[i], err = convert.StringsToTaCandle(fromCsvReverse[i])
		if err != nil {
			log.Fatalln(err)
		}
		//fmt.Println(kLength-1-i, ":", k[kLength-1-i].Array())
	}

	r, err := ta.Rsi(k, 5, 73.0, 27.0)
	rString := make([][]string, len(r))

	for i := 0; i < len(r); i++ {
		//fmt.Println(r[i].ToString())
		rString[i] = r[i].ToString()
	}

	title := [][]string{{"Time", "Close", "Result", "Signal"}}

	if err := file.ArrayToCsv("../../../history/rsi/coinmarketcap_"+coin+"_usd_k_rsi_"+strconv.Itoa(baseDay)+".csv", title); err != nil {
		log.Fatalln(err)
	}

	if err := file.ArrayAppendCsv("../../../history/rsi/coinmarketcap_"+coin+"_usd_k_rsi_"+strconv.Itoa(baseDay)+".csv", rString); err != nil {
		log.Fatalln(err)
	}
}
