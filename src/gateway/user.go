package gateway

import (
	"context"
	"dockertest-example/src/domain/model"
	"dockertest-example/src/domain/repository"
	"errors"

	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

var _ repository.UserRepository = (*User)(nil)

func NewUser(db *gorm.DB) *User {
	return &User{db}
}

func (u User) Create(ctx context.Context, user *model.User) error {
	err := u.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u User) FindByID(ctx context.Context, id int) (*model.User, error) {
	var user model.User
	if err := u.db.First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}
