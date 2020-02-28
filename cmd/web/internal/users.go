package internal

import (
	"context"
	userService "github.com/hiram66/user-service/pkg/user"
	"net/http"
)

func AddUser(ctx context.Context, user userService.User) WebResp {
	us := userService.NewService()
	result, addUserErr := us.AddUser(ctx, user)
	if addUserErr != nil {
		return WebResp{
			StatusCode: getUserErrStatusCode(addUserErr),
			JsonBody:   []byte(addUserErr.Error()),
		}
	}
	return getWebResp(result, http.StatusOK)
}

func GetUserById(ctx context.Context, id string) WebResp {
	us := userService.NewService()
	result, err := us.GetUserById(ctx, id)
	if err != nil {
		return getWebRespError(err)
	}
	return getWebResp(result, http.StatusOK)
}

func GetUser(ctx context.Context, q userService.UserQuery) WebResp {
	us := userService.NewService()
	result, err := us.GetUser(ctx, q)
	if err != nil {
		return getWebRespError(err)
	}
	return getWebResp(result, http.StatusOK)
}

func UpdateUser(ctx context.Context, id string, user userService.User) WebResp {
	us := userService.NewService()
	result, err := us.Update(ctx, id, user)
	if err != nil {
		return getWebRespError(err)
	}
	return getWebResp(result, http.StatusOK)
}

func DeleteUser(ctx context.Context, id string) WebResp {
	us := userService.NewService()
	e := us.Delete(ctx, id)
	if e != nil {
		return getWebRespError(e)
	}
	return WebResp{
		StatusCode: http.StatusNoContent,
		JsonBody:   nil,
	}
}
