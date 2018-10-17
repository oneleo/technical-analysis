package main

import (
	"fmt"
	"log"
	"time"

	"github.com/maicoin/max-exchange-api-go/models"
	"github.com/oneleo/technical-analysis/convert"
)

func main() {

	maxCandle := models.Candle{time.Unix(1519206420, int64(0)), 25000, 25000, 25000, 25000, 0.001}

	s := convert.MaxCandleToStrings(maxCandle)
	fmt.Println("String: ", s)

	c, err := convert.StringsToMaxCandle(s)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Candle: ", c)

}
