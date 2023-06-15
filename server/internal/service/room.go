package service

import (
	"github.com/yannis94/perseus/internal/core"
	"github.com/yannis94/perseus/internal/repository"
)

type RoomService struct {
    room *core.Room
    repo *repository.Repository
}

func (service *RoomService) Create() error {
    return service.repo.AddRoom(service.room)
}

func (service *RoomService) Delete() error {
    return service.repo.DeleteRoom(service.room.Id)
}
