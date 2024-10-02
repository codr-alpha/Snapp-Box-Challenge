package structs_and_constants

import "fmt"

const (
    Buffer_size             =       1000 * 1000

    R_earth                 =       6371

    Starting_fare           =       1.30
    Minimum_fare            =       3.47
    Idle_fare               =       11.90
    Moving_before5_fare     =       1.30
    Moving_after5_fare      =       0.74
)

// Point represents the data of each line of our csv input line
type Point struct {
    Id_delivery int64
    Lat float64
    Lng float64
    Timestamp float64
}

// Output_data represents the data of each line of our csv output line
type Output_data struct {
    Id_delivery int64
    Fare_estimate float64
}

// ToSlice is a method for Output_data So we can convert it to []string for writing csv
func (ot Output_data)ToSlice() []string {
    ret := make([]string, 2)
    ret[0] = fmt.Sprintf("%d", ot.Id_delivery)
    ret[1] = fmt.Sprintf("%f", ot.Fare_estimate)
    return ret
}
