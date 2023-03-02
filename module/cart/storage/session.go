package storage

import (
	"github.com/gin-contrib/sessions"
)

type sessionStore struct {
	db *sessions.Session
}

func NewSessionStore(db *sessions.Session) *sessionStore {
	return &sessionStore{db: db}
}
