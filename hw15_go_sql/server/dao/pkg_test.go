package dao

import (
	"testing"

	users "github.com/sar0868/otus_go_basic_hw/hw15_go_sql/user"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	type args struct {
		name    string
		age     int
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    *users.User
		wantErr string
	}{
		{
			name: "create user name=Name",
			args: args{
				name:    "Name",
				age:     1,
				address: "City",
			},
			want: &users.User{
				ID:      4,
				Name:    "Name",
				Age:     1,
				Address: "City",
			},
			wantErr: "",
		},
		{
			name: "error create user",
			args: args{
				name:    "Name",
				age:     0,
				address: "City",
			},
			want:    nil,
			wantErr: "age can't be less than 0",
		},
	}
	for _, tt := range tests {
		users := Users
		t.Run(tt.name, func(t *testing.T) {
			user, err := CreateUser(tt.args.name, tt.args.age, tt.args.address)
			assert.Equal(t, tt.want, user)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
			}
		})
		Users = users
	}
}

func TestGetUser(t *testing.T) {
	tests := []struct {
		name       string
		id         int
		want       *users.User
		wantResult bool
	}{
		{
			name: "get user id 1, expected user Name = Aleksey result true",
			id:   1,
			want: &users.User{
				ID:      1,
				Name:    "Aleksey",
				Age:     56,
				Address: "Tver",
			},
			wantResult: true,
		},
		{
			name:       "get user id 4, expected nil result false ",
			id:         4,
			want:       nil,
			wantResult: false,
		},
		{
			name:       "get user id -1, expected nil result false ",
			id:         -1,
			want:       nil,
			wantResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetUser(tt.id)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantResult, got1)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	type args struct {
		id      int
		name    string
		age     int
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    *users.User
		wantErr bool
	}{
		{
			name: "update user id 1, age=25",
			args: args{
				id:      1,
				name:    "Aleksey",
				age:     25,
				address: "Tver",
			},
			want: &users.User{
				ID:      1,
				Name:    "Aleksey",
				Age:     25,
				Address: "Tver",
			},
			wantErr: false,
		},
		{
			name: "update user id 4, expected false",
			args: args{
				id:      4,
				name:    "Aleksey",
				age:     25,
				address: "Tver",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "update user id 1, age -1, expected false",
			args: args{
				id:      1,
				name:    "Aleksey",
				age:     -1,
				address: "Tver",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		users := Users
		t.Run(tt.name, func(t *testing.T) {
			user, err := UpdateUser(tt.args.id, tt.args.name, tt.args.age, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, user)
		})
		Users = users
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		want    bool
		wantErr bool
	}{
		{
			name:    "delete user id 1, expected true, error false",
			id:      1,
			want:    true,
			wantErr: false,
		},
		{
			name:    "delete user id 5, expected false, error true",
			id:      5,
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteUser(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeleteUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextID(t *testing.T) {
	tests := []struct {
		name  string
		users []users.User
		want  int
	}{
		{
			name: "next id = 2",
			users: []users.User{
				{
					ID:      1,
					Name:    "Name",
					Age:     1,
					Address: "City",
				},
			},
			want: 2,
		},
		{
			name:  "next id = 1",
			users: []users.User{},
			want:  1,
		},
		{
			name: "next id = 3",
			users: []users.User{
				{
					ID:      2,
					Name:    "Name",
					Age:     1,
					Address: "City",
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, nextID(tt.users))
		})
	}
}
