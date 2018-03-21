package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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

	defer conn.Close()

	// read request
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// method & url
			mux(conn, ln)
		}
		if ln == "" {
			// if headers are done
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	m := strings.Fields(ln)[0] // method
	u := strings.Fields(ln)[1] // url
	fmt.Println("---Method:", m)
	fmt.Println("---URL:", u)

	if m == "GET" && u == "/" {
		index(conn)
	}
	if m == "GET" && u == "/about" {
		about(conn)
	}
	if m == "GET" && u == "/contact" {
		contact(conn)
	}
	if m == "GET" && u == "/apply" {
		apply(conn)
	}
	if m == "POST" && u == "/apply" {
		applyProcess(conn)
	}
}

func index(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html>
	<head>
		<title></title>
	</head>
	<body>
		<p>Hey there</p>
		<ul>
		<li><a href="/">Index</a></li>
		<li><a href="/about">About</a></li>
		<li><a href="/contact">Contact</a></li>
		<li><a href="/apply">Apply</a></li>
		</ul>
	</body>
	</html>
	`
	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprint(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html \r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html>
	<head>
		<title></title>
	</head>
	<body>
		<p>About</p>
		<ul>
		<li><a href="/">Index</a></li>
		<li><a href="/about">About</a></li>
		<li><a href="/contact">Contact</a></li>
		<li><a href="/apply">Apply</a></li>
		</ul>
	</body>
	</html>
	`
	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprint(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html \r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html>
	<head>
		<title></title>
	</head>
	<body>
		<p>Contact</p>
		<ul>
		<li><a href="/">Index</a></li>
		<li><a href="/about">About</a></li>
		<li><a href="/contact">Contact</a></li>
		<li><a href="/apply">Apply</a></li>
		</ul>
	</body>
	</html>
	`
	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprint(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html \r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func apply(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html>
	<head>
		<title></title>
	</head>
	<body>
		<p>Apply</p>
		<ul>
		<li><a href="/">Index</a></li>
		<li><a href="/about">About</a></li>
		<li><a href="/contact">Contact</a></li>
		<li><a href="/apply">Apply</a></li>
		</ul>
		<form method="post" action="/apply">
			<input type="submit" value="apply">
		</form>
	</body>
	</html>
	`
	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprint(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html \r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func applyProcess(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html>
	<head>
		<title></title>
	</head>
	<body>
		<p>Apply Process (POST request sent)</p>
		<ul>
		<li><a href="/">Index</a></li>
		<li><a href="/about">About</a></li>
		<li><a href="/contact">Contact</a></li>
		<li><a href="/apply">Apply</a></li>
		</ul>
		<form method="post" action="/apply">
			<input type="submit" value="apply">
		</form>
	</body>
	</html>
	`
	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprint(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html \r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
