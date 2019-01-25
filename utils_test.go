package main

import (
	"reflect"
	"testing"
)

func TestSplitLine(t *testing.T) {

	testTable := []struct {
		in  string
		out []string
	}{
		{"A|B|C", []string{"A", "B", "C"}},
		{"A|B|C,D", []string{"A", "B", "C,D"}},
		//Invalid inputs should return nil
		{"A|BC", nil},
		{"A|B|C|", nil},
	}

	for _, testData := range testTable {
		line, _ := splitLine(testData.in)
		if !reflect.DeepEqual(line, testData.out) {
			t.Errorf("Split line test failed. Input: %v, Expected: %v, Output: %v", testData.in, testData.out, line)
			continue
		}
		t.Logf("Split line test successful. Input: %v, Expected: %v, Output: %v", testData.in, testData.out, line)
	}

}
