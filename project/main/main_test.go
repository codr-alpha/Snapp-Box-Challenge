package main

import (
	"testing"
	"my_mod/project/structs_and_constants"
)

func equal(p1, p2 structs_and_constants.Point) bool {
	if p1.Id_delivery != p2.Id_delivery {
		return false
	}
	if p1.Lat != p2.Lat {
		return false
	}
	if p1.Lng != p2.Lng {
		return false
	}
	if p1.Timestamp != p2.Timestamp {
		return false
	}
	return true
}

func Test_toPoint(t *testing.T) {
	cases := []struct {
        in []string
        expected structs_and_constants.Point
    }{
        {[]string {"44", "4.253642", "0.345752", "100.495835"}, structs_and_constants.Point{44, 4.253642, 0.345752, 100.495835}},
        {[]string {"0", "0", "0", "0"}, structs_and_constants.Point{0, 0, 0, 0}},
        {[]string {"574820393874653", "345.45687345786", "-1", "0.1111111"}, structs_and_constants.Point{574820393874653, 345.45687345786, -1, 0.1111111}},
    }

    for _, c := range cases {
        result, err := toPoint(c.in)
        if err != nil {
        	t.Errorf("Expected no error but got %v", err)
        } else if !equal(result, c.expected) {
        	t.Errorf("Expected %v but got %v", c.expected, result)
    	}
    }
}
