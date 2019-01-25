package main

import (
	"strings"
	"sync"
)

var sem = &sync.Mutex{}

// isDependency takes a given package and looks to see if it's listed in any other packages dependency tree
func isDependency(pkg string) bool {
	for _, depList := range pkgIndex {
		for _, dep := range depList {
			if dep == pkg {
				return true
			}
		}
	}
	return false
}

// fetchEntry handles the QUERY input
func fetchEntry(pkgInfo []string) string {
	_, ok := pkgIndex[pkgInfo[0]]
	if ok {
		return "OK\n"
	}
	return "FAIL\n"
}

// editEntry handles the INDEX input
func editEntry(pkgInfo []string) string {
	pkgName := pkgInfo[0]
	deps := []string{}

	// Make sure all dependencies in pkg list already
	// but only if dep list has data
	if pkgInfo[1] != "" {
		deps = strings.Split(pkgInfo[1], ",")
		for _, dep := range deps {
			_, ok := pkgIndex[dep]
			if !ok {
				// A dependency not found in the index
				return "FAIL\n"
			}
		}
	}

	// Deps are in the index already, add new pkg entry and return
	pkgIndex[pkgName] = deps
	return "OK\n"
}

// removeEntry handles the REMOVE input
func removeEntry(pkgInfo []string) string {
	// Check if pkg exists in tree, delete if not a dependency
	_, ok := pkgIndex[pkgInfo[0]]
	if ok {
		if !isDependency(pkgInfo[0]) {
			delete(pkgIndex, pkgInfo[0])
		} else {
			// Pkg was found, but it's a dep for another package
			return "FAIL\n"
		}
	}
	// Pkg deleted successfully or wasn't found in list
	return "OK\n"
}

//The crud function will take the input given and add/remove/update/delete pkg info as necessary
func crud(input []string) string {

	ret := ""

	// Determine which function we should call with the given input
	switch input[0] {
	case "QUERY":
		ret = fetchEntry(input[1:])
	case "INDEX":
		sem.Lock()
		ret = editEntry(input[1:])
		sem.Unlock()
	case "REMOVE":
		sem.Lock()
		ret = removeEntry(input[1:])
		sem.Unlock()
	default:
		ret = "ERROR\n"
	}

	return ret
}
