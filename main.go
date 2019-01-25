package main

import (
	"bufio"
	"flag"
	"log"
	"net"
)

var pkgIndex = map[string][]string{}

// The handle function reads client input until the newline and submits it to the crud function
func handle(connection net.Conn) {
	//Read client input line-by-line (scanner.Scan() looks for \n automatically)
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		splitLine, err := validateAndSplitLine(scanner.Text())
		if err != nil {
			log.Println("[ERROR] " + err.Error())
			connection.Write([]byte("ERROR\n"))
			continue
		}

		response := crud(splitLine)
		connection.Write([]byte(response))
	}
}

func main() {

	port := flag.String("port", "8080", "The port to listen on")
	flag.Parse()

	//Create listener on given port
	log.Println("[INFO] Listening on port " + *port)
	listen, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalln("[ERROR] " + err.Error())
	}

	//Accept connections infinitely and handle them concurrently
	for {
		connection, err := listen.Accept()
		if err != nil {
			log.Println("[ERROR] " + err.Error())
		}
		go handle(connection)
	}

}
