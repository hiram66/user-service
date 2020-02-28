package mongodb

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/hiram66/user-service/storage/contracts"
	"github.com/hiram66/user-service/storage/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type userStorage struct {
	users *mongo.Collection
}

func NewUserStorage() contracts.User {
	return userStorage{users: users}
}

func (u userStorage) Add(ctx context.Context, user entities.User) (string, error) {
	id := ""
	status, e := u.users.InsertOne(ctx, u.normalizedUser(user))
	if e != nil {
		return "", e
	}
	if status.InsertedID != nil {
		s, ok := status.InsertedID.(string)
		if ok {
			id = s
		}
	}
	return id, e
}

func (u userStorage) Delete(ctx context.Context, id string) error {
	result, e := u.users.DeleteOne(ctx, bson.M{"_id": id})
	if e != nil {
		return e
	}
	if result.DeletedCount == 0 {
		return errors.New(fmt.Sprintf("user with id %s npt found", id))
	}
	return nil
}

func (u userStorage) Update(ctx context.Context, user entities.User) error {
	_, e := u.users.UpdateOne(ctx, bson.M{"_id": user.Id}, bson.M{"$set": user.GetChanges()})
	return e
}

func (u userStorage) GetById(ctx context.Context, id string) (*entities.User, error) {
	user := new(entities.User)
	find := bson.M{"_id": id}
	one := u.users.FindOne(ctx, find)
	if one.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}
	return user, one.Decode(user)
}

func (u userStorage) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	user := new(entities.User)
	find := bson.M{"email": email}
	one := u.users.FindOne(ctx, find)
	if one.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}
	return user, one.Decode(user)
}

func (u userStorage) FindByPhone(ctx context.Context, phone string) (*entities.User, error) {
	user := new(entities.User)
	find := bson.M{"phone": phone}
	one := u.users.FindOne(ctx, find)
	if one.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}
	return user, one.Decode(user)
}

func (u userStorage) FindByPhoneAndEmail(ctx context.Context, email string, phone string) (*entities.User, error) {
	user := new(entities.User)
	find := bson.M{"email": email, "phone": phone}
	one := u.users.FindOne(ctx, find)
	if one.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}
	return user, one.Decode(user)
}

func (u userStorage) genId() string {
	return uuid.New().String()
}

func (u userStorage) normalizedUser(user entities.User) entities.User {
	user.Id = u.genId()
	user.NormalizedName = u.normalizedName(user.Name, user.Family)
	return user
}

func (u userStorage) normalizedName(name, family string) string {
	return strings.ToUpper(fmt.Sprintf("%s %s", name, family))
}
