package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sar0868/otus_go_basic_hw/hw16_docker/internal/repository"
)

type ParamUser struct {
	Param string
	Value string
}

func Users(ctx context.Context, repo repository.Querier) ([]*repository.ShopUser, error) {
	users, err := repo.Users(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByParam(ctx context.Context, repo repository.Querier, params []ParamUser) (*repository.ShopUser, error) {
	if len(params) == 1 {
		switch params[0].Param {
		case "id":
			id, err := strconv.Atoi(params[0].Value)
			if err != nil {
				return nil, fmt.Errorf("error convert string to int: %w", err)
			}
			user, errGetUser := GetUserByID(ctx, repo, id)
			if errGetUser != nil {
				return nil, fmt.Errorf("error: %w", errGetUser)
			}
			return user, nil
		case "name":
			name := params[0].Value
			user, err := GetUserByName(ctx, repo, name)
			if err != nil {
				return nil, fmt.Errorf("error: %w", err)
			}
			return user, nil
		}
	}
	return nil, fmt.Errorf("request incorrect")
}

func AddUser(ctx context.Context, repo repository.Querier, newUser repository.UserAddParams) (*repository.ShopUser, error) { //nolint: lll
	id, err := repo.UserAdd(ctx, newUser)
	if err != nil {
		return nil, fmt.Errorf("don't create user: %w", err)
	}
	user, errGetUser := GetUserByID(ctx, repo, id)
	if errGetUser != nil {
		return nil, fmt.Errorf("don't found data for id=%d: %w", id, err)
	}
	return user, nil
}

func DeleteUser(ctx context.Context, repo repository.Querier, name string) error {
	_, errID := GetUserByName(ctx, repo, name)
	if errID != nil {
		return fmt.Errorf("don't found user:%w", errID)
	}
	err := repo.UserDelete(ctx, name)
	if err != nil {
		return fmt.Errorf("error delete user: %w", err)
	}
	return nil
}

func GetUserByID(ctx context.Context, repo repository.Querier, id int) (*repository.ShopUser, error) {
	user, err := repo.UserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("don't found user by id= %d", id)
	}
	return user, nil
}

func GetUserByName(ctx context.Context, repo repository.Querier, name string) (*repository.ShopUser, error) {
	user, err := repo.UserGetByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("don't found user username= %v", name)
	}
	return user, nil
}

func UserUpdate(ctx context.Context, repo repository.Querier, updateUser repository.UserUpdateParams) (*repository.ShopUser, error) { //nolint: lll
	err := repo.UserUpdate(ctx, updateUser)
	if err != nil {
		return nil, err
	}
	user, errUpdateName := GetUserByName(ctx, repo, updateUser.Name_2)
	if errUpdateName != nil {
		return nil, errUpdateName
	}
	return user, nil
}

func UsersStatistic(ctx context.Context, repo repository.Querier) ([]*repository.UsersStatisticRow, error) { //nolint: lll
	users, err := repo.UsersStatistic(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
