package calculations

import (
	"my_mod/project/structs_and_constants"
	"math"
	"time"
)

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
}

func velocity(dTime, dDistance float64) float64 {
	if (dTime <= 0) {
		return 100 + 1
	}

	return dDistance * (60 * 60) / dTime // km/h
}


func isBefore5am(t time.Time) bool {
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()

	return (hour * 60 * 60 + minute * 60 + second) <= 5 * 60 * 60 
}

func calculate(p1, p2 structs_and_constants.Point) float64 {
	dTime := p2.Timestamp - p1.Timestamp // second
	dDistance := haversine(p1, p2) // killometer
	velocity := velocity(dTime, dDistance) // km/h

	switch{
		case velocity > 100:
			return -1
		case velocity > 10:
			timestamp := int64(p1.Timestamp)
			parsedTime := time.Unix(timestamp, 0)

			if isBefore5am(parsedTime) {
				return structs_and_constants.Moving_before5_fare * dDistance
			} else {
				return structs_and_constants.Moving_after5_fare * dDistance
			}
		default:
			return structs_and_constants.Idle_fare * dTime
	}
}