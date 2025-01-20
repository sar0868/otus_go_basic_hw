package dao

import (
	"errors"

	"github.com/sar0868/otus_go_basic_hw/hw13_http/user"
)

var Users = []user.User{
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

func CreateUser(name string, age int, address string) (bool, error) {
	if age <= 0 {
		return false, errors.New("age can't be less than 0")
	}
	newUser := user.User{ID: nextID(Users), Name: name, Age: age, Address: address}
	Users = append(Users, newUser)
	return true, nil
}

func GetUser(id int) (user.User, bool) {
	for _, user := range Users {
		if user.ID == id {
			return user, true
		}
	}
	return user.User{}, false
}

func UpdateUser(id int, name string, age int, address string) (bool, error) {
	user, status := GetUser(id)
	if !status {
		return false, errors.New("don't find user")
	}
	if err := user.SetAge(age); err != nil {
		return false, err
	}
	user.SetName(name)
	user.SetAddress(address)
	return true, nil
}

func DeleteUser(id int) (bool, error) {
	ind := -1
	for i, user := range Users {
		if user.ID == id {
			ind = i
			break
		}
	}
	if ind == -1 {
		return false, errors.New("don't find user")
	}
	Users = append(Users[:ind], Users[ind+1:]...)
	return true, nil
}

func nextID(users []user.User) int {
	lastID := 0
	for _, user := range users {
		if user.ID > lastID {
			lastID = user.ID
		}
	}
	return lastID + 1
}
