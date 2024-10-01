package main

import (
	"my_mod/project/calculations"
	"my_mod/project/structs_and_constants"
	"encoding/csv"
	"os"
	"strconv"
	"fmt"
	"sync"
)

func strToFloat64(s string) float64 {
	floatValue, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}

	return floatValue
}

func strToInt64(s string) int64 {
	intValue, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return intValue
}

func toPoint(s []string) structs_and_constants.Point {
	return structs_and_constants.Point{
					Id_delivery: strToInt64(s[0]),
					Lat: strToFloat64(s[1]),
					Lng: strToFloat64(s[2]),
					Timestamp: strToFloat64(s[3]),
				}
}

func main() {
	file, err := os.Open("../../input/sample_data.csv")
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()

	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan structs_and_constants.Point)
	go calculations.Process(ch, &wg)

	reader := csv.NewReader(file)

	if _, err := reader.Read(); err != nil && err.Error() != "EOF" {
		panic(err)
	}

	for {
		record, err := reader.Read()
		if err != nil && err.Error() != "EOF" {
			panic(err)
			break
		}
		if len(record) == 0 {
			break
		}
		ch <- toPoint(record)
	}
	close(ch)

	wg.Wait()

	fmt.Println("successfully done!!!")
}
