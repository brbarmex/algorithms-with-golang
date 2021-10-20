package arrays

import (
	"testing"
	"reflect"
)

func TestReverseArray(t *testing.T){
	var cases = []struct {
		arr []int32
		expected []int32
		}{
			{
				arr: []int32{6676,3216,4063,8373,423,586,8850,6762},
				expected: []int32{6762,8850,586,423,8373,4063,3216,6676},
			},
			{
				arr: []int32{1,4,3,2},
				expected: []int32{2,3,4,1},
			},
			{
				arr: []int32{1,2,3,4,5},
				expected: []int32{5,4,3,2,1},
			},
		}

	for _, scenario := range cases {
		got := Reverse(scenario.arr)
		if !reflect.DeepEqual(got, scenario.expected) {
			t.Errorf("ArrayReverse(arr) = %d; want %d", got,scenario.expected)
		}
	}
}