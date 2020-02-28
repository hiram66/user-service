package internal

import (
	"fmt"
	"github.com/hiram66/user-service/pkg/errors"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type WebResp struct {
	StatusCode int
	JsonBody   []byte
}

func GetJsonError(err error) []byte {
	return []byte(fmt.Sprintf("{\"error\" : \"%s\"}", err.Error()))
}

func getUserErrStatusCode(err error) int {
	switch err.(type) {
	case errors.UserRequiredFieldErr:
		return http.StatusBadRequest
	case errors.UserDuplicateFieldError:
		return http.StatusNotAcceptable
	case errors.UserNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func getWebResp(data interface{}, successCode int) WebResp {
	jsonResult, err := jsoniter.Marshal(&data)
	if err != nil {
		return getWebRespError(err)
	}
	return WebResp{
		StatusCode: successCode,
		JsonBody:   jsonResult,
	}
}

func getWebRespError(err error) WebResp {
	return WebResp{
		StatusCode: getUserErrStatusCode(err),
		JsonBody:   []byte(err.Error()),
	}
}
