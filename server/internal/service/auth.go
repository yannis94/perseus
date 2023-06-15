package service

import (
	"time"

	"github.com/yannis94/perseus/internal/config"
)

type Auth struct {

}

func NewAuthService() *Auth {
    return &Auth{}
}

func (auth *Auth) Login(user User) (User, error) {
    jwt_exp := time.Now().Add(config.JWT_EXP * time.Minute)
    refresh_exp := time.Now().Add(config.REFRESH_EXP * time.Hour)

    jwt, err := createToken(user.id, jwt_exp)
    
    if err != nil {
        return User{}, err
    }

    refresh_token, err := createToken(user.id, refresh_exp)

    if err != nil {
        return User{}, err
    }

    user.token = jwt
    user.refresh_token = refresh_token
    return User{}, nil
}

func (auth *Auth) Logout(user_id int) error {
    //delete user session in db
    return nil
}

func createToken(user_id int, exp_time time.Time) (string, error) {
    //build jwt token 
    return "super-tkn", nil
}

func (auth *Auth) VerifyToken(token string) error {
    //verify token 
    return nil
}

func (auth *Auth) RefreshToken(refresh_token string) (string, error) {
    // refresh token if refresh_token is still valid
    return "new-token", nil
}
