package main

import (
	"fmt"
	"reflect"
	"testing"
)

// Write your test here
func TestFindOddNumber(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedOutput int
	}{
		{input: []int{7}, expectedOutput: 7},
		{input: []int{0}, expectedOutput: 0},
		{input: []int{1, 1, 2}, expectedOutput: 2},
		{input: []int{0, 1, 0, 1, 0}, expectedOutput: 0},
		{input: []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}, expectedOutput: 4},
	}

	for _, testCase := range testCases {
		output := FindOddNumber(testCase.input)
		count := 0
		for _, num := range testCase.input {
			if num == output {
				count++
			}
		}
		if !reflect.DeepEqual(output, testCase.expectedOutput) {
			t.Errorf("%v should return %v, but It occurs %v times (which is odd).\n", testCase.input, testCase.expectedOutput, count)
		} else {
			fmt.Printf("%v should return %v, because it occurs %v times (which is odd). Success\n", testCase.input, testCase.expectedOutput, count)
		}
	}
}
