package modells

import "time"

type User struct {
	ID        uint64    `json:"id, omitempty"`
	Name      string    `json:"name, omitempty"`
	NickName  string    `json:"nickName, omitempty"`
	Email     string    `json:"Email, omitempty"`
	Password  string    `json:"Password, omitempty"`
	CreatedAt time.Time `json:"CreatedAt, omitempty"`
}
