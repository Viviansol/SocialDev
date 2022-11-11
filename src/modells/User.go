package modells

import (
	"api/src/security"
	"errors"
	"github.com/badoux/checkmail"
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

func (user *User) PrepareUser(stage string) error {
	if err := user.validateUser(stage); err != nil {
		return err
	}

	if err := user.formatUser(stage); err != nil {
		return err
	}
	return nil
}

func (user *User) validateUser(stage string) error {
	if user.Name == "" {
		return errors.New("Name field is mandatory")
	}
	if user.NickName == "" {
		return errors.New("Nickname field is mandatory")
	}
	if user.Email == "" {
		return errors.New("Email field is mandatory")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Invalid Email Format")
	}
	if stage == " registration" && user.Password == "" {
		return errors.New("Password field is mandatory")
	}

	return nil
}

func (user *User) formatUser(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.NickName = strings.TrimSpace(user.NickName)
	user.Email = strings.TrimSpace(user.Email)
	if stage == "registration" {
		passwordWithHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(passwordWithHash)
	}

	return nil
}
