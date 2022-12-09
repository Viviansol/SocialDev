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
	statement, err := u.db.Prepare("INSERT INTO users(name, nickName, email, password)VALUES(?,?,?,?)")
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
		"SELECT id, name, nickName, email, createdAt FROM users WHERE name LIKE ? or nickName LIKE ? ",
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
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u user) GetUserById(ID uint64) (modells.User, error) {
	rows, err := u.db.Query(
		"SELECT id, name, nickname, email, createdAt FROM users WHERE id = ?", ID)
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

func (u user) DeleteUser(Id uint64) error {
	statement, err := u.db.Prepare("DELETE * FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(Id); err != nil {
		return err
	}
	return nil
}

func (u user) GetUserByEmail(email string) (modells.User, error) {
	rows, err := u.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if err != nil {
		return modells.User{}, err
	}

	defer rows.Close()
	var user modells.User
	if rows.Next() {
		if err = rows.Scan(&user.ID, &user.Password); err != nil {
			return modells.User{}, err
		}
	}
	return user, nil

}

func (u user) FollowUser(userId, followerId uint64) error {
	statement, err := u.db.Prepare("INSERT INTO followers(user_id, follower_id)VAlUES(?,?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil

}

func (u user) UnfollowUser(userId, followerId uint64) error {
	statement, err := u.db.Prepare("delete from followers where user_id = ? and  follower_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil

}

func (u user) SearchFollowers(id uint64) ([]modells.User, error) {
	rows, err := u.db.Query(`select u.id, u.name, u.nickName, u.email, u.createdAt
									from users u inner join followers on u.id = s.followers_id where s.user_id =?`, id)

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
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)

	}
	return users, nil

}

func (u user) SearchFollowing(id uint64) ([]modells.User, error) {
	rows, err := u.db.Query(`select u.id, u.name, u.nickName, u.email, u.createdAt
									from users u inner join followers on u.id = s.followers_id where s.follower_id =?`, id)

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
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)

	}
	return users, nil

}

func (u user) SearchPasswordById(id uint64) (string, error) {

	row, err := u.db.Query("SELECT password FROM users WHERE id = ?", id)

	if err != nil {
		return " ", err
	}

	defer row.Close()

	var user modells.User
	if row.Next() {
		if err = row.Scan(
			&user.Password,
		); err != nil {
			return " ", err
		}

	}

	return user.Password, nil
}

func (u user) UpdatePassword(userID uint64, password string) error {
	statement, err := u.db.Prepare("UPDATE users SET password = ? where id =?")

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(password, userID); err != nil {
		return err
	}

	return nil
}
