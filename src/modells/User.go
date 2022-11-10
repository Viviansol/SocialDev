package modells

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id, omitempty"`
	Name      string    `json:"name, omitempty"`
	NickName  string    `json:"nickName, omitempty"`
	Email     string    `json:"Email, omitempty"`
	Password  string    `json:"Password, omitempty"`
	CreatedAt time.Time `json:"CreatedAt, omitempty"`
}

func (user *User) PrepareUser() error {
	if err := user.validateUser(); err != nil {
		return err
	}

	user.formatUser()
	return nil
}

func (user *User) validateUser() error {
	if user.Name == "" {
		return errors.New("Name field is mandatory")
	}
	if user.NickName == "" {
		return errors.New("Nickname field is mandatory")
	}
	if user.Email == "" {
		return errors.New("Email field is mandatory")
	}
	if user.Password == "" {
		return errors.New("Password field is mandatory")
	}

	return nil
}

func (user *User) formatUser() {
	user.Name = strings.TrimSpace(user.Name)
	user.NickName = strings.TrimSpace(user.NickName)
	user.Email = strings.TrimSpace(user.Email)

}
