package main

import (
	"fmt"

	maxapi "github.com/maicoin/max-exchange-api-go"
	mymax "github.com/oneleo/technical-analysis/api/max"
)

func main() {

	client := maxapi.NewClient()
	defer client.Close()
	ethtwdTicker, _ := mymax.GetTicker(client, "ethtwd")
	fmt.Println(ethtwdTicker)

}
