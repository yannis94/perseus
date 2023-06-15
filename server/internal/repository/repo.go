package repository

import (
	"github.com/yannis94/perseus/internal/core"
)

type Database interface {
    AddUser(*core.User) error
    AddRoom(*core.Room) error
    AddMessage(*core.Message) error
    GetUserByEmail(email string) (*core.User, error)
    GetUserById(id int) (*core.User, error)
    GetRooms() ([]*core.Room, error)
    GetMessages(room_id int) ([]*core.Message, error)
    DeleteRoom(room_id int) error
}

type Repository struct {
    db Database
}

func NewRepository(db Database) *Repository {
    return &Repository{ db: db }
}

func (repo *Repository) AddUser(user *core.User) error {
    return repo.db.AddUser(user)
}

func (repo *Repository) AddRoom(room *core.Room) error {
    return repo.db.AddRoom(room)
}

func (repo *Repository) AddMessage(msg *core.Message) error {
    return repo.db.AddMessage(msg)
}

func (repo *Repository) GetUserById(id int) (*core.User, error) {
    return repo.db.GetUserById(id)
}

func (repo *Repository) GetUserByEmail(email string) (*core.User, error) {
    return repo.db.GetUserByEmail(email)
}

func (repo *Repository) GetRooms() ([]*core.Room, error) {
    return repo.db.GetRooms()
}

func (repo *Repository) GetMessages(room_id int) ([]*core.Message, error) {
    return repo.db.GetMessages(room_id)
}

func (repo *Repository) DeleteRoom(id int) error {
    return repo.db.DeleteRoom(id)
}
