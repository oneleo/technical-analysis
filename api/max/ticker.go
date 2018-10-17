package max

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	maxapi "github.com/maicoin/max-exchange-api-go"
	"github.com/maicoin/max-exchange-api-go/models"
)

// GetTicker method will get ticker of specific market.
// Response Class (Status 200):
// ticker is within 24 hours, models.Ticker.At is timestamp in seconds since Unix epoch
// Reference: https://max.maicoin.com/documents/api_list
// https://mholt.github.io/json-to-go/
func GetTicker(client maxapi.API, market string) (ticker models.Ticker, err error) {

	// Ticker returns a ticker of specific market.
	//
	// Available `CallOption`:
	// (None)
	// Reference:
	// https://github.com/maicoin/max-exchange-api-go/blob/master/interface.go
	results, err := client.Ticker(context.Background(), market)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return models.Ticker{}, err
	}
	//m := map[string]string{}

	//json.Unmarshal([]byte(resultBytes), &m)
	//fmt.Println(m)
	//fmt.Println(string(resultBytes))
	resultBytes, err := json.MarshalIndent(results, "", "\t")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return models.Ticker{}, err
	}

	json.Unmarshal(resultBytes, &ticker)
	return ticker, err
}
