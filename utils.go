package main

import (
	"errors"
	"index/suffixarray"
	"strings"
)

// The splitLine function breaks apart the client input and returns it to be acted upon
func splitLine(pkgText string) ([]string, error) {

	// Find the indices of the "|" character in the input string. Bail if we don't have 2 of them.
	index := suffixarray.New([]byte(pkgText))
	pipes := index.Lookup([]byte("|"), -1)
	if len(pipes) != 2 {
		return nil, errors.New("Invalid client input " + pkgText)
	}

	split := strings.Split(pkgText, "|")
	return split, nil
}
