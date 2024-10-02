package calculations

import (
	"my_mod/project/structs_and_constants"
	"math"
	"time"
	"encoding/csv"
	"os"
	"sync"
	"fmt"
)

const (
	minuteToSecond		=		60
	hourToSecond		=		60 * 60
)

// haversine method gets 2 Point and return the haversine distance of that 2 point
func haversine(p1, p2 structs_and_constants.Point) float64 {
	lat1 := p1.Lat * math.Pi / 180
	lng1 := p1.Lng * math.Pi / 180
	lat2 := p2.Lat * math.Pi / 180
	lng2 := p2.Lng * math.Pi / 180

	dLat := lat2 - lat1
	dLng := lng2 - lng1

	a := math.Sin(dLat / 2) * math.Sin(dLat / 2) +
		math.Cos(lat1) * math.Cos(lat2) * math.Sin(dLng / 2) * math.Sin(dLng / 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return structs_and_constants.R_earth * c
} // test with unit tests in calculations_test.go

func velocity(dTime, dDistance float64) float64 {
	if (dTime <= 0) {
		return 100 + 1
	}
	return dDistance * hourToSecond / dTime // km/h
} // test with unit tests in calculations_test.go

// isBefore5am checks if a time.Time is between (00:00, 05:00] or not
// it calculates how many seconds we are after 00:00 and checks if it is bigger than 0 and not bigger that (05:00)
func isBefore5am(t time.Time) bool {
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()

	v := (hour * hourToSecond + minute * minuteToSecond + second)

	return v <= 5 * hourToSecond && v > 0
} // test with unit tests in calculations_test.go

// calculateFare gets 2 Point and calculate its Fare
// if it is an invalid segment it returns -1
func calculateFare(p1, p2 structs_and_constants.Point) float64 {
	dTime := p2.Timestamp - p1.Timestamp // second
	dDistance := haversine(p1, p2) // killometer
	velocity := velocity(dTime, dDistance) // km/h

	switch {
		case velocity > 100:
			return -1
		case velocity > 10:
			timestamp := int64(p1.Timestamp)
			parsedTime := time.Unix(timestamp, 0)

			// according to pdf it checks the p1 time is before 5 am or not and use the correct fare based on that
			if isBefore5am(parsedTime) {
				return structs_and_constants.Moving_before5_fare * dDistance
			} else {
				return structs_and_constants.Moving_after5_fare * dDistance
			}
		default:
			return structs_and_constants.Idle_fare * dTime / hourToSecond
	}
} // we test this with E2E tests

/*
Process gets Point one by one from ch channel and process them using calculateFare function
and send the data it wants to write as Output_data type to ch2 for writingToCSV function to write it
it use _wg to tell main that it is finished
and
use wg to make sure writingToCSV do its job before Process
done is for tellling main it couldn't finish the job without error
done2 is for writingToCSV to tell Process it couldn't finish the job without error
*/
func Process(ch chan structs_and_constants.Point, _wg *sync.WaitGroup, done chan struct{}) {
	defer _wg.Done()

	var wg sync.WaitGroup
	wg.Add(1)
	done2 := make(chan struct{})

	ch2 := make(chan structs_and_constants.Output_data, structs_and_constants.Buffer_size)
	go writingToCSV(ch2, &wg, done2) // making the thread to write to csv file

	p1 := structs_and_constants.Point{}
	ot := structs_and_constants.Output_data{}

	for p2 := range ch {
		if p1.Id_delivery != p2.Id_delivery {
			// making sure fare is at least Minimum_fare
			if structs_and_constants.Minimum_fare > ot.Fare_estimate {
				ot.Fare_estimate = structs_and_constants.Minimum_fare
			}
			// check if p1 represent a real Point or not
			if p1.Id_delivery != 0 {
				select {
					case <-done2:
						// if writingToCSV have an error we close done to tell main we have an error and return
						close(done)
						return
					default:
						// otherwise we add ot to ch2
						ch2 <- ot
				}
			}

			// the new point become p1 we set Id_delivery and Fare_estimate
			ot.Id_delivery = p2.Id_delivery
			ot.Fare_estimate = structs_and_constants.Starting_fare
			p1 = p2
		} else {
			// for p1 and p2 if the id of both of them is the same we calculate cost using calculateFare function
			// and if p2 is invalid we do nothing
			// otherwise we add cost to out Fare_estimate for the id
			cost := calculateFare(p1, p2)
			if cost != -1 {
				ot.Fare_estimate += cost
				p1 = p2
			}
		}
	}

	// making sure fare is at least Minimum_fare
	if structs_and_constants.Minimum_fare > ot.Fare_estimate {
		ot.Fare_estimate = structs_and_constants.Minimum_fare
	}
	// check if p1 represent a real Point or not
	if p1.Id_delivery != 0 {
		ch2 <- ot
	}
	close(ch2) // close ch2 so writingToCSV can finish its job

	wg.Wait()
	close(done)
} // we test this with E2E tests

/*
writingToCSV creates the csv ouput file
writing it line by line
based on the Ouput_data from ch channel
done is for tellling Process it couldn't finish the job without error
wg is for telling Process that it finished its job
*/
func writingToCSV(ch chan structs_and_constants.Output_data, wg *sync.WaitGroup, done chan struct{}) {
	defer wg.Done()
	file, err := os.Create("../../output/output.csv")
	if err != nil {
		fmt.Println("Can not create output file!!!")
		fmt.Println(err)
		close(done)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	// Ensure all buffered data is written to the file when the function exits
	defer writer.Flush()

	// writing headers and handling error
	head := []string{"id_delivery", "fare_estimate"}
	if err := writer.Write(head); err != nil {
		fmt.Println("Can not write to file!!!")
		fmt.Println(err)
		close(done)
		return
	}

	for v := range ch {
		// write v and handle error
		if err := writer.Write(v.ToSlice()); err != nil {
			fmt.Println("Can not write to file!!!")
			fmt.Println(err)
			close(done) // for telling Process it didn't do the job without error
			return
		}
	}
	close(done)
} // we test this with E2E tests