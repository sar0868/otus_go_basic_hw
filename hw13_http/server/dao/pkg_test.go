package dao

import (
	"testing"

	"github.com/sar0868/otus_go_basic_hw/hw13_http/user"
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
		want    bool
		wantErr string
	}{
		{
			name: "create user name=Name",
			args: args{
				name:    "Name",
				age:     1,
				address: "City",
			},
			want:    true,
			wantErr: "",
		},
		{
			name: "error create user",
			args: args{
				name:    "Name",
				age:     0,
				address: "City",
			},
			want:    false,
			wantErr: "age can't be less than 0",
		},
	}
	for _, tt := range tests {
		users := Users
		t.Run(tt.name, func(t *testing.T) {
			result, err := CreateUser(tt.args.name, tt.args.age, tt.args.address)
			assert.Equal(t, tt.want, result)
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
		want       user.User
		wantName   string
		wantResult bool
	}{
		{
			name: "get user id 1, expected user Name = Aleksey result true",
			id:   1,
			want: user.User{
				ID:      1,
				Name:    "Aleksey",
				Age:     56,
				Address: "Tver",
			},
			wantName:   "Aleksey",
			wantResult: true,
		},
		{
			name:       "get user id 4, expected User{} result false ",
			id:         4,
			want:       user.User{},
			wantName:   "",
			wantResult: false,
		},
		{
			name:       "get user id -1, expected User{} result false ",
			id:         -1,
			want:       user.User{},
			wantName:   "",
			wantResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetUser(tt.id)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantName, got.Name)
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
		want    bool
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
			want:    true,
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
			want:    false,
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
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		users := Users
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdateUser(tt.args.id, tt.args.name, tt.args.age, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateUser() = %v, want %v", got, tt.want)
			}
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
		users []user.User
		want  int
	}{
		{
			name: "next id = 2",
			users: []user.User{
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
			users: []user.User{},
			want:  1,
		},
		{
			name: "next id = 3",
			users: []user.User{
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
