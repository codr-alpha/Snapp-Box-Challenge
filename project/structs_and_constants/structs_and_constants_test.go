package structs_and_constants

import (
	"testing"
	"reflect"
)

func TestOutput_data_ToSlice(t *testing.T) {
	cases := []struct {
        a int64
        b float64
        expected []string
    }{
        {0, 0, []string {"0", "0.000000"}},
        {-3432432, 1.02347983249, []string{"-3432432", "1.023480"}},
        {15847395748295, 3.141592654123, []string{"15847395748295", "3.141593"}},
    }

    for _, c := range cases {
    	output_data := &Output_data{c.a, c.b}
        result := output_data.ToSlice()
        if !reflect.DeepEqual(result, c.expected) {
        	t.Errorf("Expected %v but got %v", c.expected, result)
    	}
    }
}