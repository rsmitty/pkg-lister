package main

import (
	"reflect"
	"testing"
)

func TestValidateAndSplitLine(t *testing.T) {

	testTable := []struct {
		in  string
		out []string
	}{
		{"A|B|C", []string{"A", "B", "C"}},
		{"A|B|C,D", []string{"A", "B", "C,D"}},
		//Invalid inputs should return nil
		{"A|BC", nil},           //Too few pipes
		{"A|B|C|", nil},         //Too many pipes
		{"A|B|C,D,1234,🙀", nil}, //How did this crying cat get in here?
		{"A|B|C D", nil},        //Space in package list

	}

	for _, testData := range testTable {
		line, _ := validateAndSplitLine(testData.in)
		if !reflect.DeepEqual(line, testData.out) {
			t.Errorf("validateAndSplitLine line test failed. Input: %v, Expected: %v, Output: %v", testData.in, testData.out, line)
			continue
		}
		t.Logf("validateAndSplitLine test successful. Input: %v, Expected: %v, Output: %v", testData.in, testData.out, line)
	}

}
