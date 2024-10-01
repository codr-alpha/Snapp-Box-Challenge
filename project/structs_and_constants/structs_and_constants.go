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


type Point struct {
    Id_delivery int64
    Lat float64
    Lng float64
    Timestamp float64
}

type Output_data struct {
    Id_delivery int64
    Fare_estimate float64
}

func (ot Output_data)ToSlice() []string {
    rt := make([]string, 2)
    rt[0] = fmt.Sprintf("%d", ot.Id_delivery)
    rt[1] = fmt.Sprintf("%f", ot.Fare_estimate)
    return rt
}
