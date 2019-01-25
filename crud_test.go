package main

import (
	"testing"
)

func TestIsDependency(t *testing.T) {
	pkgIndex = map[string][]string{"A": {"B", "C"}, "B": {}, "C": {}}

	testTable := []struct {
		in  string
		out bool
	}{
		{"A", false},
		{"B", true},
	}

	for _, testData := range testTable {
		depTrue := isDependency(testData.in)
		if depTrue != testData.out {
			t.Errorf("Dependency test failed. Input: %v, Expected: %v, Output: %v", testData.in, testData.out, depTrue)
			continue
		}
		t.Logf("Dependency test succeeded. Input: %v, Expected: %v, Output: %v", testData.in, testData.out, depTrue)
	}
}

func TestFetchEntry(t *testing.T) {
	pkgIndex = map[string][]string{"A": {"B", "C"}, "B": {}, "C": {}}

	testTable := []struct {
		in  []string
		out string
	}{
		{[]string{"A", "B,C"}, "OK\n"},   //A is present in test data
		{[]string{"X", "Y,Z"}, "FAIL\n"}, //X isn't in our index
	}

	for _, testData := range testTable {
		retString := fetchEntry(testData.in)
		if retString != testData.out {
			t.Errorf("fetchEntry test failed. Input: %v, Expected: %v, Output: %v", testData.in, testData.out, retString)
			continue
		}
		t.Logf("fetchEntry test succeeded. Input: %v, Expected: %v, Output: %v", testData.in, testData.out, retString)
	}
}

func TestEditEntry(t *testing.T) {
	pkgIndex = map[string][]string{"A": {"B", "C"}, "B": {}, "C": {}}

	testTable := []struct {
		in  []string
		out string
	}{
		{[]string{"D", ""}, "OK\n"},          //Add D with no deps
		{[]string{"X", "B,C"}, "OK\n"},       //Adds X with deps
		{[]string{"A", "B,C,D"}, "OK\n"},     //Updates A
		{[]string{"E", "B,C,D,F"}, "FAIL\n"}, //Attempts to add E, but F won't be in dep list
	}

	for _, testData := range testTable {
		retString := editEntry(testData.in)
		if retString != testData.out {
			t.Errorf("editEntry test failed. Input: %v, Expected: %v, Output: %v", testData.in, testData.out, retString)
			continue
		}
		t.Logf("editEntry test succeeded. Input: %v, Expected: %v, Output: %v", testData.in, testData.out, retString)
	}
}

func TestRemoveEntry(t *testing.T) {
	pkgIndex = map[string][]string{"A": {"B", "C"}, "B": {}, "C": {}}

	testTable := []struct {
		in  []string
		out string
	}{
		{[]string{"B", ""}, "FAIL\n"}, //B is a dependency for A
		{[]string{"Z", ""}, "OK\n"},   //Nonexistent
		{[]string{"A", ""}, "OK\n"},   //A isn't a dependency

	}

	for _, testData := range testTable {
		retString := removeEntry(testData.in)
		if retString != testData.out {
			t.Errorf("removeEntry test failed. Input: %v, Expected: %v, Output: %v", testData.in, testData.out, retString)
			continue
		}
		t.Logf("removeEntry test succeeded. Input: %v, Expected: %v, Output: %v", testData.in, testData.out, retString)
	}
}

// This function should test all CRUD capabilities, so the tests above may be superfluous
func TestCrud(t *testing.T) {
	pkgIndex = map[string][]string{"A": {"B", "C"}, "B": {}, "C": {}}

	testTable := []struct {
		in  []string
		out string
	}{
		{[]string{"QUERY", "A", ""}, "OK\n"},   //Present
		{[]string{"QUERY", "Z", ""}, "FAIL\n"}, //Not present

	}

	for _, testData := range testTable {
		retString := crud(testData.in)
		if retString != testData.out {
			t.Errorf("crud test failed. Input: %v, Expected: %v, Output: %v", testData.in, testData.out, retString)
			continue
		}
		t.Logf("crud test succeeded. Input: %v, Expected: %v, Output: %v", testData.in, testData.out, retString)
	}
}
