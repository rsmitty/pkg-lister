package main

import (
	"bufio"
	"net"
	"testing"
)

func TestHandle(t *testing.T) {
	pkgIndex = map[string][]string{"A": {"B", "C"}, "B": {}, "C": {}}

	testTable := []struct {
		in  string
		out string
	}{
		//Lookup tests against handler
		{"QUERY|A|\n", "OK\n"},   //Exist
		{"QUERY|Z|\n", "FAIL\n"}, //Doesn't exist

		//Index tests against handler
		{"INDEX|D|\n", "OK\n"},          //Add no deps
		{"INDEX|X|B,C\n", "OK\n"},       //Add with deps
		{"INDEX|A|B,C,D\n", "OK\n"},     //Update A
		{"INDEX|E|B,C,D,F\n", "FAIL\n"}, //Attempts to add E, but F won't be in dep list

		//Remove tests against handler
		{"REMOVE|B|\n", "FAIL\n"}, //Is dep
		{"REMOVE|Z|\n", "OK\n"},   //Nonexistent
		{"REMOVE|A|\n", "OK\n"},   //Not a dependency

		//Bogus input
		{"REMOVE|A||\n", "ERROR\n"},         //Not a dependency
		{"NOTAREALOPTION|A|\n", "ERROR\n"},  //Not a dependency
		{"INDEX|B|C,D,1234,🙀\n", "ERROR\n"}, //How did this crying cat get in here?
		{"INDEX|B|C D\n", "ERROR\n"},        //Space in package list
	}

	//net.Pipe returns both ends of a client/server pair
	client, server := net.Pipe()
	defer client.Close()
	defer server.Close()

	//Pass mock server to handle function in the background
	//This allows us to read from the client as data is returned
	go func(net.Conn) {
		handle(server)
	}(server)

	//Write data to client side and read response
	scanner := bufio.NewScanner(client)
	for _, testData := range testTable {
		client.Write([]byte(testData.in))
		scanner.Scan()
		if scanner.Text()+"\n" != testData.out {
			t.Errorf("handle test failed. Input: %v, Expected: %v, Output: %v", testData.in, testData.out, scanner.Text())
			continue
		}
		t.Logf("handle test succeeded. Input: %v, Expected: %v, Output: %v", testData.in, testData.out, scanner.Text())
	}

}
