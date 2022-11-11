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

func (u user) GetUserById(ID uint64) (modells.User, error) {
	rows, err := u.db.Query(
		"SELECT id, name, nickaname, email, createdAt FROM users WHERE id = ?", ID)
	if err != nil {
		return modells.User{}, err
	}
	defer rows.Close()
	var person modells.User
	if rows.Next() {
		if err = rows.Scan(
			&person.ID,
			&person.Name,
			&person.NickName,
			&person.CreatedAt,
		); err != nil {
			return modells.User{}, err
		}

	}
	return person, nil
}

func (u user) UpdateUser(Id uint64, user modells.User) error {
	statement, err := u.db.Prepare("update user set name = ?, nickname = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(user.Name, user.NickName, user.Email, Id); err != nil {
		return err
	}
	return nil
}
