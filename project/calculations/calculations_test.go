package calculations

import (
	"time"
	"math"
	"testing"
	"my_mod/project/structs_and_constants"
)

func Test_haversine(t *testing.T) {
	cases := []struct {
        p1 structs_and_constants.Point
        p2 structs_and_constants.Point
        expected float64
    }{
        {structs_and_constants.Point{0, 37.7749, -122.4194, 0}, structs_and_constants.Point{0, 37.7749, -122.4194, 0}, 0},
        {structs_and_constants.Point{0, 0, 0, 0}, structs_and_constants.Point{0, 90, 0, 0}, 10000},
        {structs_and_constants.Point{0, 34.0522, -118.2437, 0}, structs_and_constants.Point{0, 34.0522, -74.0060, 0}, 3944},
        {structs_and_constants.Point{0, 0, 179, 0}, structs_and_constants.Point{0, 0, -179, 0}, 222},
    }

    for _, c := range cases {
        result := haversine(c.p1, c.p2)
        if math.Abs(c.expected - result) > 100 {
        	t.Errorf("Expected %v but got %v", c.expected, result)
    	}
    }
}

func Test_velocity(t *testing.T) {
	cases := []struct {
        time, distance, expected float64
    }{
    	{0, 10, 101},
    	{100, 0, 0},
    	{7200, 500, 250},
    }

    for _, c := range cases {
        result := velocity(c.time, c.distance)
        if c.expected != result {
        	t.Errorf("Expected %v but got %v", c.expected, result)
    	}
    }
}

func Test_isBefore5am(t *testing.T) {
	cases := []struct {
		time time.Time
		expected bool
	}{
		{time.Date(2024, 10, 2, 0, 0, 0, 0, time.UTC), false},
		{time.Date(2024, 10, 2, 0, 0, 1, 0, time.UTC), true},
		{time.Date(2024, 10, 2, 5, 0, 0, 0, time.UTC), true},
		{time.Date(2024, 10, 2, 5, 0, 1, 0, time.UTC), false},
	}

	for _, c := range cases {
        result := isBefore5am(c.time)
        if c.expected != result {
        	t.Errorf("Expected %v but got %v", c.expected, result)
    	}
    }
}