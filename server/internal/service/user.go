package service

import (
	"github.com/yannis94/perseus/internal/core"
	"github.com/yannis94/perseus/internal/repository"
)

type UserService struct {
    user *core.User
    repo *repository.Repository
}

func NewUserService(db *repository.Repository) *UserService {
    return &UserService{
        repo: db,
    }
}

func (service *UserService) Create() error {
    return service.repo.AddUser(service.user)
}

func (service *UserService) GetUserByEmail(email string) (*core.User, error) {
    return service.repo.GetUserByEmail(email)
}

func (service *UserService) AddMessage(msg *core.Message) error {
    return service.repo.AddMessage(msg)
}
