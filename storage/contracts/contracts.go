package contracts

import (
	"context"
	"github.com/hiram66/user-service/storage/entities"
)

type User interface {
	Add(ctx context.Context, user entities.User) (string, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, user entities.User) error
	GetById(ctx context.Context, id string) (*entities.User, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	FindByPhone(ctx context.Context, phone string) (*entities.User, error)
	FindByPhoneAndEmail(ctx context.Context, email string, phone string) (*entities.User, error)
}
