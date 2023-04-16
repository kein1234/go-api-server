package repository

import "sample-api/domain/model"

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindByID(id int64) (*model.User, error)
	Store(user *model.User) error
	Update(user *model.User) error
	Delete(id int64) error
}
