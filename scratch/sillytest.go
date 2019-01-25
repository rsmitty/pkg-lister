package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "INDEX|wuta|\n")

	fmt.Fprintf(conn, "INDEX|wutb|\n")

	fmt.Fprintf(conn, "INDEX|wutc|\n")

	fmt.Fprintf(conn, "INDEX|wut|wuta,wutb,wutc\n")

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
