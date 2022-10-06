package repository

import (
	"api/src/modells"
	"database/sql"
)

type user struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *user {
	return &user{db}
}

func (u user) CreateUser(user modells.User) (uint64, error) {
	return 0, nil
}
