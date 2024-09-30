package structs_and_constants


const (
    Buffer_size     =       1000 * 1000
    R_earth         =       6371
)


type Point struct {
    id_delivery int64
    lat float64
    lng float64
    timestamp int64
}

type Output_data struct {
    id_delivery int64
    fare_estimate float64
}
