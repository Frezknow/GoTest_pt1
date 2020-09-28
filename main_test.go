package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{22, 21}, 43},
		test{[]int{21, 21, 10}, 52},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if mySum(v.data...) != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}
