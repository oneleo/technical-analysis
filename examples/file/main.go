package main

import (
	"fmt"
	"log"

	"github.com/oneleo/technical-analysis/file"
)

func main() {

	title := [][]string{{"Time", "Open", "High", "Low", "Close", "Volume"}}

	if err := file.ArrayToCsv("../../history/Example_K.csv", title); err != nil {
		log.Fatalln(err)
	}

	k1 := []string{"1519206420", "25000", "25000", "25000", "25000", "0.001"}

	k2 := []string{"1529206420", "15120", "15199", "15120", "15199", "1"}

	var data [][]string
	data = append(data, k1)
	data = append(data, k2)

	if err := file.ArrayAppendCsv("../../history/Example_K.csv", data); err != nil {
		log.Fatalln(err)
	}

	fromCsv, err := file.CsvToArray("../../history/Example_K.csv")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("From CSV file:\n", fromCsv)

}
