package usecase

import (
	"fmt"

	"enigmacamp.com/be-lms-university/model/entity"
	"enigmacamp.com/be-lms-university/repository"
)

type UserUseCase interface {
	FindById(id string) (entity.User, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func (u *userUseCase) FindById(id string) (entity.User, error) {
	// menambahkan validasi custom dipersilahkan
	// misalnya membuat pesan yang lebih informatif
	// ID tidak ditemukan
	// business logic

	user, err := u.repo.Get(id)
	if err != nil {
		return entity.User{}, fmt.Errorf("user with ID %s not found", id)
	}

	return user, nil
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
