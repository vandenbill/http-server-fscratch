package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHttp(t *testing.T) {
	url := "http://localhost:8080"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("fail send req :", err)
		return
	}

	req.Header.Set("User-Agent", "MyApp")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("fail get res :", err)
		return
	}
	defer resp.Body.Close()

	body := make([]byte, 10<<10)
	n, err := resp.Body.Read(body)

	fmt.Println("status :", resp.Status)
	fmt.Println("protocol :", resp.Proto)
	fmt.Println("res : ", string(body[:n]))
}
