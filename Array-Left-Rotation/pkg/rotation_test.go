package rotation

import (
	"testing"
	"reflect"
)

func TestLeftRotation(t *testing.T)  {
	
	var cases = []struct {
		d int32
		arr []int32
		expected []int32
	}{
		{
			d: int32(7),
			arr: []int32{98,67,35,1,74,79,7,26,54,63,24,17,32,81},
			expected : []int32{26,54,63,24,17,32,81,98,67,35,1,74,79,7},
		},
		{
			d: int32(5),
			arr: []int32{1,2,3,4,5},
			expected : []int32{5,1,2,3,4},
		},
	}

	for _, scenarios := range cases {
		got := RotateLeft(scenarios.d,scenarios.arr)
		if !reflect.DeepEqual(got, scenarios.expected) {
			t.Errorf("RotateLeft() = %d; Want %d;",got, scenarios.expected)
		}
	}
}