package user

import (
	"context"
	"fmt"
	usersError "github.com/hiram66/user-service/pkg/errors"
	"github.com/hiram66/user-service/storage"
	"github.com/hiram66/user-service/storage/entities"
	"strings"
)

type userService struct{}

func NewService() Service {
	return userService{}
}

func (u userService) AddUser(ctx context.Context, user User) (AddStatus, error) {
	validationErr := u.checkMissingFields(user)
	if validationErr != nil {
		return AddStatus{}, validationErr
	}
	duplicateFieldsErr := u.checkDuplicateFields(ctx, user.Email, user.Phone)
	if duplicateFieldsErr != nil {
		return AddStatus{}, duplicateFieldsErr
	}
	insertedId, err := storage.UserStorage.Add(ctx, u.toEntity(user))
	return AddStatus{Id: insertedId}, err
}

func (u userService) GetUserById(ctx context.Context, id string) (UserResponse, error) {
	user, e := storage.UserStorage.GetById(ctx, id)
	if e != nil {
		return UserResponse{}, e
	}
	if user == nil {
		return UserResponse{}, usersError.UserNotFound{
			GenericError: usersError.GenericError{Err: fmt.Sprintf("user with id %s not found", id)},
		}
	}
	return u.entityToResp(*user), nil
}

func (u userService) GetUser(ctx context.Context, q UserQuery) (UserResponse, error) {
	user, e := u.doGet(ctx, q)
	if e != nil {
		return UserResponse{}, e
	}
	return u.entityToResp(user), nil
}

func (u userService) Update(ctx context.Context, id string, user User) (UserResponse, error) {
	err := u.checkDuplicateFields(ctx, user.Email, user.Phone)
	if err != nil {
		return UserResponse{}, err
	}
	oldUser, err := storage.UserStorage.GetById(ctx, id)
	if err != nil {
		return UserResponse{}, err
	}
	if oldUser == nil {
		return UserResponse{}, usersError.UserNotFound{
			GenericError: usersError.GenericError{Err: fmt.Sprintf("user with id %s not found", id)},
		}
	}
	newUser := u.mergeWithEntity(user, *oldUser)
	err = storage.UserStorage.Update(ctx, newUser)
	if err != nil {
		return UserResponse{}, err
	}
	return u.entityToResp(newUser), nil
}

func (u userService) Delete(ctx context.Context, id string) error {
	return storage.UserStorage.Delete(ctx, id)
}

func (u userService) toEntity(user User) entities.User {
	return entities.User{
		Email:  user.Email,
		Name:   user.Name,
		Family: user.Family,
		Phone:  user.Phone,
	}
}

func (u userService) checkMissingFields(user User) error {
	fields := make([]string, 0)
	if len(user.Phone) == 0 {
		fields = append(fields, "phone")
	}
	if len(user.Name) == 0 {
		fields = append(fields, "name")
	}
	if len(user.Family) == 0 {
		fields = append(fields, "family")
	}
	if len(user.Email) == 0 {
		fields = append(fields, "email")
	}
	if len(fields) != 0 {
		return usersError.UserRequiredFieldErr{
			Fields:  fields,
			Message: "missing required fields",
		}
	}
	return nil
}

func (u userService) checkDuplicateFields(ctx context.Context, email, phone string) error {
	email = strings.TrimSpace(email)
	phone = strings.TrimSpace(phone)
	if len(email) != 0 {
		user, e := storage.UserStorage.FindByEmail(ctx, email)
		if e != nil {
			return e
		}
		if user != nil {
			return usersError.UserDuplicateFieldError{Field: "email", Value: email}
		}
	}

	if len(phone) != 0 {
		user, e := storage.UserStorage.FindByPhone(ctx, phone)
		if e != nil {
			return e
		}
		if user != nil {
			return usersError.UserDuplicateFieldError{Field: "phone", Value: phone}
		}
	}
	return nil
}

func (u userService) doGet(ctx context.Context, query UserQuery) (entities.User, error) {
	if len(query.Email) == 0 && len(query.Phone) == 0 {
		return entities.User{}, usersError.UserNotFound{GenericError: usersError.GenericError{
			Err: fmt.Sprintf("phone or email should be specified")}}
	}

	if len(query.Email) == 0 {
		user, e := storage.UserStorage.FindByPhone(ctx, query.Phone)
		if e != nil {
			return entities.User{}, e
		}
		if user == nil {
			return entities.User{}, usersError.UserNotFound{GenericError: usersError.GenericError{
				Err: fmt.Sprintf("user with phone %s not found", query.Phone)}}
		}
		return *user, nil
	}

	if len(query.Phone) == 0 {
		user, e := storage.UserStorage.FindByEmail(ctx, query.Email)
		if e != nil {
			return entities.User{}, e
		}
		if user == nil {
			return entities.User{}, usersError.UserNotFound{GenericError: usersError.GenericError{
				Err: fmt.Sprintf("user with email %s not found", query.Email)}}
		}
		return *user, nil
	}

	user, e := storage.UserStorage.FindByPhoneAndEmail(ctx, query.Email, query.Phone)
	if e != nil {
		return entities.User{}, e
	}
	if user == nil {
		return entities.User{}, usersError.UserNotFound{GenericError: usersError.GenericError{
			Err: fmt.Sprintf("user with email %s and phone %s not found", query.Email, query.Phone)}}
	}
	return *user, nil
}

func (u userService) entityToResp(user entities.User) UserResponse {
	return UserResponse{
		User: User{
			Name:   user.Name,
			Family: user.Family,
			Phone:  user.Phone,
			Email:  user.Email,
		},
		Id: user.Id,
	}
}

func (u userService) mergeWithEntity(user User, entity entities.User) entities.User {
	userEntity := &entity
	userEntity.SetName(user.Name)
	userEntity.SetEmail(user.Email)
	userEntity.SetFamily(user.Family)
	userEntity.SetPhone(user.Phone)
	return *userEntity
}
