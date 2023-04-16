package interactor

import (
	"sample-api/domain/model"
	"sample-api/domain/repository"
)

type UserInteractor struct {
	UserRepository repository.UserRepository
}

func (i *UserInteractor) FindAll() ([]*model.User, error) {
	return i.UserRepository.FindAll()
}

func (i *UserInteractor) FindByID(id int64) (*model.User, error) {
	return i.UserRepository.FindByID(id)
}

func (i *UserInteractor) Store(user *model.User) error {
	return i.UserRepository.Store(user)
}

func (i *UserInteractor) Update(user *model.User) error {
	return i.UserRepository.Update(user)
}

func (i *UserInteractor) Delete(id int64) error {
	return i.UserRepository.Delete(id)
}
