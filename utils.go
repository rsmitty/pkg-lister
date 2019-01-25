package main

import (
	"errors"
	"index/suffixarray"
	"regexp"
	"strings"
)

// The validateAndSplitLine function breaks apart the client input and returns it to be acted upon
func validateAndSplitLine(pkgText string) ([]string, error) {

	// Find the indices of the "|" character in the input string. Bail if we don't have 2 of them.
	index := suffixarray.New([]byte(pkgText))
	pipes := index.Lookup([]byte("|"), -1)
	if len(pipes) != 2 {
		return nil, errors.New("Invalid client input with pipes " + pkgText)
	}
	split := strings.Split(pkgText, "|")

	//Regex must match empty string (^$) or be in the set A-Z, a-z, 0-9, comma, underscore
	validPkgString, _ := regexp.MatchString("^$|^[A-Za-z0-9,_]+$", split[2])
	if !validPkgString {
		return nil, errors.New("Invalid client input on pkg list " + pkgText)
	}

	return split, nil
}
