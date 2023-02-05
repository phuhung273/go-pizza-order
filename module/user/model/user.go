package model

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

var (
	ErrUsernameCannotBeEmpty = errors.New("username cannot be empty")
	ErrPasswordCannotBeEmpty = errors.New("password cannot be empty")
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex; size:50"`
	Password string `gorm:"size:200"`
}

func (i *User) Validate() error {
	i.Username = strings.TrimSpace(i.Username)

	if i.Username == "" {
		return ErrUsernameCannotBeEmpty
	} else if i.Password == "" {
		return ErrPasswordCannotBeEmpty
	}

	return nil
}
