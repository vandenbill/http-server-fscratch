package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("server start...")

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			return
		}

		go func() {
			defer conn.Close()

			fmt.Printf("Req coming from | %s \n", conn.RemoteAddr().String())

			s := bufio.NewScanner(conn)
			for !s.Scan() {
			}

			req := s.Text()
			if !isHTTPRequest(string(req)) {
				fmt.Println("invalid request")
				return
			}

			fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprint(conn, "Content-Type: text/html\r\n\r\n")
			fmt.Fprint(conn, "<p>"+strings.Split(req, " ")[0]+"<p>")
		}()
	}
}

func isHTTPRequest(data string) bool {
	validMethods := []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS", "PATCH"}
	for _, method := range validMethods {
		if len(data) > len(method) && data[:len(method)] == method {
			return true
		}
	}
	return false
}
