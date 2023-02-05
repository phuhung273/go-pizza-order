package model

import (
	"errors"
	"strings"
)

var (
	ErrUsernameCannotBeEmpty = errors.New("title cannot be empty")
	ErrPasswordCannotBeEmpty = errors.New("password cannot be empty")
)

type RegisterRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (i *RegisterRequest) Validate() error {
	i.Username = strings.TrimSpace(i.Username)

	if i.Username == "" {
		return ErrUsernameCannotBeEmpty
	} else if i.Password == "" {
		return ErrPasswordCannotBeEmpty
	}

	return nil
}
