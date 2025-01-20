package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/sar0868/otus_go_basic_hw/hw13_http/user"
)

var users = []user.User{
	{
		ID:      1,
		Name:    "Aleksey",
		Age:     56,
		Address: "Tver",
	},
	{
		ID:      2,
		Name:    "Irina",
		Age:     60,
		Address: "Tver",
	},
	{
		ID:      3,
		Name:    "Maria",
		Age:     27,
		Address: "Tver",
	},
}

func main() {
	var ADDRESS string
	var PORT string
	flag.StringVar(&ADDRESS, "address", "127.0.0.1", "- address for server")
	flag.StringVar(&PORT, "port", "8080", "- port for server")
	flag.Parse()

	fmt.Printf("server run: %s:%s\n", ADDRESS, PORT)

	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/user", getUser)
	// if err := http.ListenAndServe(ADDRESS+":"+PORT, nil); err != nil {
	// 	fmt.Println("Error run server:", err)
	// }
	server := &http.Server{
		Addr:              ADDRESS + ":" + PORT,
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       10 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello World")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Println(r.Context())
	w.Header().Set("Content-Type", "application/json")
}
