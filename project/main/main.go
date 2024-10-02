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

// strToFloat64 get a string and convert it to a float64 type and return error if it couldn't do it
func strToFloat64(s string) (float64, error) {
	floatValue, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return floatValue, nil
}

// strToInt64 get a string and convert it to a int64 type and return error if it couldn't do it
func strToInt64(s string) (int64, error) {
	intValue, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}

// toPoint convert a []string to Point type and return an error if converting some string returns an error
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

/*
main opens the csv input file
reading it line by line
and converts each line with toPoint function to a Point
and sends it to a channel that another thread process it
it uses sync.WaitGroup for making sure other thread finish its job before main thread
*/
func main() {
	file, err := os.Open("../../input/sample_data.csv")
	if err != nil {
		fmt.Println("Can not open file!!!")
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)

	// reading first line (headers) and handling error
	if _, err := reader.Read(); err != nil && err.Error() != "EOF" {
		fmt.Println("Can not read file!!!")
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	done := make(chan struct{}) // a channel for calculations.Process to tell main that it did't finish its job successfuly
	ch := make(chan structs_and_constants.Point, structs_and_constants.Buffer_size)
	go calculations.Process(ch, &wg, done) // making the thread to process Points

	for {
		record, err := reader.Read()
		// handling error of reading file
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Can not read file!!!")
			fmt.Println(err)
			close(ch)
			return
		}
		// making sure the record is not empty
		if len(record) == 0 {
			break
		}
		/*
			if the calculations.Process finish its job unsuccessfuly we end our job in here
			otherwise we convert the line we read to a Point using toPoint
			handle the error of toPoint
			and send the Point to the channel 
		*/
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
	close(ch) // close the channel so calculations.Process can finish its job

	wg.Wait() // it will block until calculations.Process finish its job

	fmt.Println("Successfully done!!!")
}
