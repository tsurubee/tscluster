package main

import (
	"log"
	"github.com/tsurubee/tscluster"
	"fmt"
	"os"
	"encoding/csv"
	"strconv"
)

func main() {
	f, err := os.Open("./sample.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var dataset [][]float64
	dataLen := len(rawCSVData) - 1
	dataNum := len(rawCSVData[0]) - 1
	for i := 1; i < dataNum + 1; i++ {
		var data []float64
		for j := 1; j < dataLen + 1; j++ {
			value, err := strconv.ParseFloat(rawCSVData[j][i], 64)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, value)
		}
		dataset = append(dataset, data)
	}

	tc := tscluster.NewTscluster(tscluster.DTW)
	labels, err := tc.Kmedoids(dataset, 3, 20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(labels)
}
