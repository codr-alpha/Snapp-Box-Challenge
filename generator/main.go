package main

import (
	"encoding/csv"
	"os"
	"fmt"
	"math/rand"
)

func randint(l, r int) int64 {
	return int64(rand.Intn(r - l + 1) + l)
}

func newData(id_delivery int64, lat, lng float64, timestamp int64) (int64, float64, float64, int64) {
	if id_delivery == 0 || randint(1, 40) == 40 {
		id_delivery += 1
		timestamp = randint(1500000000, 1700000000)
		lat = rand.Float64() * 180 - 90
		lng = rand.Float64() * 360 - 180
		return id_delivery, lat, lng, timestamp
	}
	timeInc := randint(20, 60)
	timestamp += timeInc

	deltaLat := rand.Float64() * (0.004 * 2) - (0.004)
	lat += deltaLat

	deltaLng := rand.Float64() * (0.004 * 2) - (0.004)
	lng += deltaLng

	return id_delivery, lat, lng, timestamp
}

func main() {
	var i, n int64
	fmt.Scanf("%d %d", &i, &n)
	rand.Seed(i)

	output_name := fmt.Sprintf("input%d.csv", i)
	output_address := fmt.Sprintf("../input/%s", output_name)
	file, err := os.Create(output_address)
	if err != nil {
		fmt.Println("Can not create output file!!!")
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	// Ensure all buffered data is written to the file when the function exits
	defer writer.Flush()

	// writing headers and handling error
	head := []string{"id_delivery", "lat", "lng", "timestamp"}
	if err := writer.Write(head); err != nil {
		fmt.Println("Can not write to file!!!")
		fmt.Println(err)
		return
	}

	var id_delivery, timestamp int64
	var lat, lng float64
	for i := int64(0); i < n; i++ {
		id_delivery, lat, lng, timestamp = newData(id_delivery, lat, lng, timestamp)
		ot := make([]string, 4)
		ot[0] = fmt.Sprintf("%d", id_delivery)
		ot[1] = fmt.Sprintf("%f", lat)
		ot[2] = fmt.Sprintf("%f", lng)
		ot[3] = fmt.Sprintf("%d", timestamp)

		if err := writer.Write(ot); err != nil {
			fmt.Println("Can not write to file!!!")
			fmt.Println(err)
			return
		}
	}
}