package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sar0868/otus_go_basic_hw/hw15_go_sql/user"
)

func main() {
	var IP string
	var PORT string
	flag.StringVar(&IP, "ip", "127.0.0.1", "-ip")
	flag.StringVar(&PORT, "port", "8080", "-port")
	flag.Parse()

	getUsersURL := fmt.Sprintf("http://%s:%s/users", IP, PORT)
	getUsers(getUsersURL)

	getUserURLID2 := fmt.Sprintf("http://%s:%s/user?id=2", IP, PORT)
	getUser(getUserURLID2)

	postAddUserURL := fmt.Sprintf("http://%s:%s/add_user", IP, PORT)
	newUser := users.User{Name: "User", Age: 1, Address: "City"}
	postAddUser(postAddUserURL, newUser)
}

func getUsers(url string) {
	resp, err := http.Get(url) //nolint
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error HTTP-response: %d\n", resp.StatusCode)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error read", err)
		return
	}

	var users []users.User
	errUnmarshal := json.Unmarshal(body, &users)
	if errUnmarshal != nil {
		fmt.Println("Error deserialization", err)
		return
	}
	fmt.Println("Users:\n", users)
}

func getUser(url string) {
	resp, err := http.Get(url) //nolint
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error HTTP-response: %d\n", resp.StatusCode)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error read", err)
		return
	}
	var user users.User
	errUnmarshal := json.Unmarshal(body, &user)
	if errUnmarshal != nil {
		fmt.Println("Error deserialization", err)
		return
	}
	fmt.Println("User:\n", user)
}

func postAddUser(url string, newUser users.User) {
	data, errMarshal := json.Marshal(newUser)
	if errMarshal != nil {
		fmt.Println("Error serialization", errMarshal)
		return
	}
	resp, err := http.Post(url, //nolint
		"application/json", strings.NewReader(string(data)))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error HTTP-response: %d\n", resp.StatusCode)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error read", err)
		return
	}
	var user users.User
	errUnmarshal := json.Unmarshal(body, &user)
	if errUnmarshal != nil {
		fmt.Println("Error deserialization", err)
		return
	}
	fmt.Println("New user:\n", user)
}
