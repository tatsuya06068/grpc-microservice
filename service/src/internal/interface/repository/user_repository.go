package repository

import (
    "fmt"
    "github.com/myusername/myservice/internal/entity"
)

type userRepository struct {
    users map[string]*entity.User
}

func NewUserRepository() *userRepository {
    return &userRepository{
        users: map[string]*entity.User{
            "1": {ID: "1", Name: "John Doe", Age: 30},
        },
    }
}

func (r *userRepository) FindByID(id string) (*entity.User, error) {
    user, exists := r.users[id]
    if !exists {
        return nil, fmt.Errorf("user not found")
    }
    return user, nil
}