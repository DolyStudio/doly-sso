package business

import (
	"time"
)

type UserToken struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Thumnail string `json:"thumnail"`

	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	TokenExpiry  time.Time `json:"tokenExpiry"`
}

type Opts func(userToken *UserToken) error

func WithInfo(email, name, thumnail string) Opts {
	return func(userToken *UserToken) error {
		userToken.Email = email
		userToken.Name = name
		userToken.Thumnail = thumnail
		return nil
	}
}

func WithToken(accessToken, refreshToken string) Opts {
	return func(userToken *UserToken) error {
		userToken.AccessToken = accessToken
		userToken.RefreshToken = refreshToken
		return nil
	}
}

func WithTokenExpired(t time.Time) Opts {
	return func(userToken *UserToken) error {
		userToken.TokenExpiry = t
		return nil
	}
}

func NewUser(opts ...Opts) (UserToken, error) {
	o := &UserToken{}

	for _, opt := range opts {
		if err := opt(o); err != nil {
			return UserToken{}, err
		}
	}

	return *o, nil
}
