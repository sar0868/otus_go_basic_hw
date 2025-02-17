package user

import "errors"

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetAge(age int) error {
	if age <= 0 {
		return errors.New("age can't be less than 0")
	}
	u.Age = age
	return nil
}

func (u *User) SetAddress(address string) {
	u.Address = address
}
