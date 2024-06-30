package usecase

import "github.com/myusername/myservice/internal/entity"

type IUserUsecase interface {
	GetUser(id string) (*entity.User, error)
}

type userUsecase struct {
	userRepository IUserRepository
}

func NewUserUsecase(repo IUserRepository) IUserUsecase {
	return &userUsecase{userRepository: repo}
}

func (u *userUsecase) GetUser(id string) (*entity.User, error) {
	return u.userRepository.FindByID(id)
}

type IUserRepository interface {
	FindByID(id string) (*entity.User, error)
}
