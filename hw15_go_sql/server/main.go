package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/sar0868/otus_go_basic_hw/hw13_http/server/dao"
	"github.com/sar0868/otus_go_basic_hw/hw13_http/user"
)

func main() {
	var IP string
	var PORT string
	flag.StringVar(&IP, "ip", "127.0.0.1", "- ip for server")
	flag.StringVar(&PORT, "port", "8080", "- port for server")
	flag.Parse()

	fmt.Printf("server run: %s:%s\n", IP, PORT)

	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/user", getUser)
	http.HandleFunc("/add_user", createUser)
	server := &http.Server{
		Addr:              IP + ":" + PORT,
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       10 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("Request received get data users")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dao.Users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("request method not equal GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	idStr := r.Form.Get("id")
	fmt.Printf("Request received get data for id %s\n", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Error convert string to int")
		return
	}
	user, result := dao.GetUser(id)
	if !result {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("request method not equal POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var newUser user.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding JSON: %v", err)
		return
	}
	user, err := dao.CreateUser(newUser.Name, newUser.Age, newUser.Address)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	fmt.Printf("Created user:\n\tid=%d\n\tname=%s\n\tage=%d\n\taddress=%s\n",
		user.ID, user.Name, user.Age, user.Address)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
