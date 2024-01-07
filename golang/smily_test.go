package main

import (
	"fmt"
	"reflect"
	"testing"
)

// Write your test here
func TestCountSmilyFace(t *testing.T) {
	testCases := []struct {
		input          []string
		expectedOutput int
	}{
		{input: []string{":)", ";(", ";}", ":-D"}, expectedOutput: 2},
		{input: []string{";D", ":-(", ":-)", ";~)"}, expectedOutput: 3},
		{input: []string{";]", ":[", ";*", ":$", ";-D"}, expectedOutput: 1},
	}

	for _, testCase := range testCases {
		output := CountSmilyFace(testCase.input)
		if !reflect.DeepEqual(output, testCase.expectedOutput) {
			t.Errorf("With input %v: Expected %v, got %v. Failure\n", testCase.input, testCase.expectedOutput, output)
		} else {
			fmt.Printf("With input %v: Your function should return %v, got %v. Success\n", testCase.input, testCase.expectedOutput, output)
		}
	}
}
