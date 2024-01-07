package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// Write your test here
func TestManipulate(t *testing.T) {
	testCases := []struct {
		input          []string
		expectedOutput []string
	}{
		{input: []string{"a"}, expectedOutput: []string{"a"}},
		{input: []string{"a", "b"}, expectedOutput: []string{"ab", "ba"}},
		{input: []string{"a", "b", "c"}, expectedOutput: []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{input: []string{"a", "a", "b", "b"}, expectedOutput: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}},
	}

	for _, testCase := range testCases {
		output := Manipulate(testCase.input)
		sort.Strings(output)
		sort.Strings(testCase.expectedOutput)
		if !reflect.DeepEqual(output, testCase.expectedOutput) {
			t.Errorf("With input '%v':\nExpected: %v\nGot: %v\n", testCase.input, testCase.expectedOutput, output)
		} else {
			fmt.Printf("With input %v: Your function should return %v, got %v. Success\n", testCase.input, testCase.expectedOutput, output)
		}
	}
}
