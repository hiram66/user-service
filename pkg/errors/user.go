package errors

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

type GenericError struct {
	Err string `json:"error"`
}

func (ge GenericError) Error() string {
	jsonError, _ := jsoniter.Marshal(&ge)
	return string(jsonError)
}

type UserRequiredFieldErr struct {
	Fields  []string `json:"fields"`
	Message string   `json:"message"`
}

func (u UserRequiredFieldErr) Error() string {
	jsonError, e := jsoniter.Marshal(&u)
	if e != nil {
		ge := GenericError{Err: e.Error()}
		return ge.Error()
	}
	return string(jsonError)
}

type UserDuplicateFieldError struct {
	Field, Value string
}

func (u UserDuplicateFieldError) Error() string {
	genericError := GenericError{Err: fmt.Sprintf("user with %s %s already exists", u.Field, u.Value)}
	return genericError.Error()
}

type UserNotFound struct {
	GenericError
}
