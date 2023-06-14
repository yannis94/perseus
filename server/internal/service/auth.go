package service

import "time"

type Auth struct {

}

func NewAuthService() *Auth {
    return &Auth{}
}

func (auth *Auth) Login(email, password string) (User, error) {
    //find user in db
    return User{}, nil
}

func (auth *Auth) Logout(user_id int) error {
    //delete user session in db
    return nil
}

func (auth *Auth) CreateToken(user_id int, exp_time time.Time) (string, error) {
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
