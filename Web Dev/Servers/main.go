package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	// make infinite loop
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		// make goroutine
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I think you say: %s\n", ln)
	}
	defer conn.Close()

	fmt.Println("Code got here")
}
