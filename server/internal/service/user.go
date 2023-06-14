package service

import (
	"time"

	"github.com/yannis94/perseus/internal/config"
)

type User struct {
    id int
    Username string
    Email string
    password string
    token string
    refresh_token string
}

type UserService struct {
    User User
    Auth AuthService
    Message MessageService
    Room RoomService
}

func NewUserService() *UserService {
    roomService := NewRoomService()
    msgService := NewMessageService()
    authService := NewAuthService()

    return &UserService{ Auth: authService, Message: msgService, Room: roomService}
}

func (u *UserService) Login(email, password string) (User, error) {
    user, err := u.Auth.Login(email, password)

    if err != nil {
        return User{}, err
    }

    jwt_exp := time.Now().Add(config.JWT_EXP * time.Minute)
    refresh_exp := time.Now().Add(config.REFRESH_EXP * time.Hour)
    jwt, err := u.Auth.CreateToken(user.id, jwt_exp)
    
    if err != nil {
        return User{}, err
    }

    refresh_token, err := u.Auth.CreateToken(user.id, refresh_exp)

    if err != nil {
        return User{}, err
    }

    user.token = jwt
    user.refresh_token = refresh_token

    return user, nil
}

func (u *UserService) TokenValidation(token string) error {
    return u.Auth.VerifyToken(token)
}

func (u *UserService) TokenRefresh(refresh_token string) (string, error) {
    return u.Auth.RefreshToken(refresh_token)
}

func (u *UserService) Logout(user_id int) error {
    return u.Auth.Logout(user_id)
}

func (u *UserService) JoinRoom(room_id int) (Room, error) {
    return u.Room.Join(room_id)
}

func (u *UserService) LeaveRoom(room_id int) error {
    return u.Room.Leave(room_id)
}

func (u *UserService) SendMessage(message_content string) error {
    return u.Message.New(u.User.id, u.User.Username, message_content)
}
