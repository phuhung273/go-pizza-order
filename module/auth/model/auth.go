package model

import (
	"errors"
	"strings"
)

var (
	ErrUsernameCannotBeEmpty = errors.New("username cannot be empty")
	ErrPasswordCannotBeEmpty = errors.New("password cannot be empty")
)

type AuthRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (i *AuthRequest) Validate() error {
	i.Username = strings.TrimSpace(i.Username)

	if i.Username == "" {
		return ErrUsernameCannotBeEmpty
	} else if i.Password == "" {
		return ErrPasswordCannotBeEmpty
	}

	return nil
}
