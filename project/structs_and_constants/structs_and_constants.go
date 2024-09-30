package structs_and_constants


const (
    Buffer_size     =       1000 * 1000
    R_earth         =       6371
)


type Point struct {
    Id_delivery int64
    Lat float64
    Lng float64
    Timestamp int64
}

type Output_data struct {
    Id_delivery int64
    Fare_estimate float64
}
