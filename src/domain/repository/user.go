package repository

import (
	"context"
	"dockertest-example/src/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindByID(ctx context.Context, id int) (*model.User, error)
}
