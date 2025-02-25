package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sar0868/otus_go_basic_hw/hw15_go_sql/internal/repository"
)

type ParamUser struct {
	Param string
	Value string
}

func Users(ctx context.Context, repo repository.Querier) ([]*repository.ShopUser, error) {
	users, err := repository.Querier.Users(repo, ctx)
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
				return nil, fmt.Errorf("error convert string to int")
			}
			user, errGetUser := repository.Querier.UserByID(repo, ctx, id)
			if errGetUser != nil {
				return nil, fmt.Errorf("don't found data for id= %d", id)
			}
			return user, nil
		case "name":
			name := params[0].Value
			user, err := repository.Querier.UserGetByName(repo, ctx, name)
			if err != nil {
				return nil, fmt.Errorf("don't found data for name=%s", name)
			}
			return user, nil
		}
	}
	return nil, fmt.Errorf("request incorrect")
}
