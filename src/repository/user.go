package repository

import (
	"api/src/modells"
	"database/sql"
	"fmt"
)

type user struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *user {
	return &user{db}
}

func (u user) CreateUser(user modells.User) (uint64, error) {
	statement, err := u.db.Prepare("INSERT INTO users(name, nick, email, password) values (?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	result, err := statement.Exec(user.Name, user.NickName, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastInsertId), nil

}

func (u user) GetUser(nameOrNick string) ([]modells.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)
	rows, err := u.db.Query(
		"SELECT id, name, email, createdAt FROM users WHERE name LIKE ? or nickName LIKE ? ",
		nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []modells.User
	for rows.Next() {
		var user modells.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			user.Email,
			user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
