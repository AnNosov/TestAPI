package store

import "TestAPI/internal/app/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(string) (*model.User, error)
}
