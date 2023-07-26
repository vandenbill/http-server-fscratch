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

			w := bufio.NewWriter(conn)
			w.WriteString(
				fmt.Sprintf("%s\r\n%s\r\n\r\n%s",
					"HTTP/1.1 200 OK",
					"Content-Type: text/html",
					strings.Split(s.Text(), " ")[0],
				))
			w.Flush()
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
