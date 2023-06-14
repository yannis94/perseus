package service

import "time"

type AuthService interface {
    Login(email, password string) (User, error)
    Logout(user_id int) error
    CreateToken(user_id int, timout time.Time) (string, error)
    VerifyToken(token string) error
    RefreshToken(token string) (string, error)
}

type MessageService interface {
    New(user_id int, username string, content string) error
}

type RoomService interface {
    Create(name string, capacity int) (Room, error)
    Join(room_id int) (Room, error)
    Leave(room_id int) error
}


