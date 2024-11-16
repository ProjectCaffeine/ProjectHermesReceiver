package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 1024)

	n, err := resp.Body.Read(data)

	if err != nil && err.Error() != "EOF"{
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Print(string(data[:n]))
}
