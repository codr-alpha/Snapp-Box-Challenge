package calculations

import (
	"my_mod/project/structs_and_constants"
	"math"
)


func Haversine(p1, p2 structs_and_constants.Point) float64 {
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