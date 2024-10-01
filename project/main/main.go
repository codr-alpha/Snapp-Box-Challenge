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

func strToFloat64(s string) (float64, error) {
	floatValue, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return floatValue, nil
}

func strToInt64(s string) (int64, error) {
	intValue, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}

func toPoint(s []string) (structs_and_constants.Point, error) {
	err := make([]error, 4)
	var id_delivery int64
	var lat, lng, timestamp float64
	id_delivery, err[0] = strToInt64(s[0])
	lat, err[1] = strToFloat64(s[1])
	lng, err[2] = strToFloat64(s[2])
	timestamp, err[3] = strToFloat64(s[3])
	for _, e := range err {
		if e != nil {
			return structs_and_constants.Point{}, e
		}
	}
	return structs_and_constants.Point{
					Id_delivery: id_delivery,
					Lat: lat,
					Lng: lng,
					Timestamp: timestamp,
				}, nil
}

func main() {
	file, err := os.Open("../../input/sample_data.csv")
	if err != nil {
		fmt.Println("Can not open file!!!")
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)

	if _, err := reader.Read(); err != nil && err.Error() != "EOF" {
		fmt.Println("Can not read file!!!")
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	done := make(chan struct{})
	ch := make(chan structs_and_constants.Point, structs_and_constants.Buffer_size)
	go calculations.Process(ch, &wg, done)

	for {
		record, err := reader.Read()
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Can not read file!!!")
			fmt.Println(err)
			close(ch)
			return
		}
		if len(record) == 0 {
			break
		}
		select {
			case <-done:
				return
			default:
				p, e := toPoint(record)
				if e != nil {
					fmt.Println(e)
					close(ch)
					return
				}
				ch <- p
		}
	}
	close(ch)

	wg.Wait()

	fmt.Println("Successfully done!!!")
}
