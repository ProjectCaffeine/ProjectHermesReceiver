package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type User struct {
	Name string
	Email string
}

func createUser(name, email string) User {
	return User {
		Name: name,
		Email: email,
	}
}

func main() {
	for {
		fmt.Print(`What action would you like to perform?
			1: GET /
			2: POST /User
			q: Exit
		`)	

		r := bufio.NewReader(os.Stdin)

		input, err := r.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.Replace(input, "\r\n", "", -1)
		input = strings.Replace(input, "\n", "", -1)

		if input == "q" {
			break
		} else if input == "1" {
			getIndex()
		} else if input == "2" {
			postToUser()
		} else {
			fmt.Printf("Invalid input: '%s'\n", input)
		}
	}
}

func postToUser() {
	user := createUser("John", "test@test.com")
	jsonUser, err := json.Marshal(user)
	r := bytes.NewReader(jsonUser)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8080/User", "json", r)

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Status: %s\n", resp.Status)
}

func getIndex() {
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
	fmt.Print("\n")
}
