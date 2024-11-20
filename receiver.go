package main

import (
	"bufio"
	"bytes"
	"encoding/json"

	//"encoding/json"
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
			2: GET /User?123
			3: POST /User
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
			getUserById()
		} else if input == "3" {
			postToUser()
		} else {
			fmt.Printf("Invalid input: '%s'\n", input)
		}
	}
}

func postDummyDataToUser() {
	data := []byte("test")

	r, err := http.NewRequest("POST", "http://localhost:8080/User", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Status: %s\n", res.Status)
}

func postToUser() {
	user := createUser("John", "test@test.com")
	jsonUser, err := json.Marshal(user)
	r := bytes.NewBuffer(jsonUser)

	if err != nil {
		log.Fatal(err)
	}


	resp, err := http.Post("http://localhost:8080/User", "application/json", r)

	if err != nil {
		log.Fatal(err)
	}
	
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n", resp.Status)
}

func getUserById() {
	resp, err := http.Get("http://localhost:8080/User?id=123")

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
