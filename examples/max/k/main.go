package main

import (
	"fmt"
	"strconv"
	"time"

	maxapi "github.com/maicoin/max-exchange-api-go"
	mymax "github.com/oneleo/technical-analysis/api/max"
)

func main() {
	client := maxapi.NewClient()
	defer client.Close()

	ethtwdK, err := mymax.GetK1Min(client, "ethtwd", int32(time.Unix(1519206420, 0).Unix()))
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Print("K line at ", strconv.FormatInt(ethtwdK.Time.Unix(), 10), ": \n", ethtwdK)

}
