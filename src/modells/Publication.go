package modells

import (
	"errors"
	"strings"
	"time"
)

type Publication struct {
	Id         uint64    `json:"id, omitempty"`
	Title      string    `json:"title, omitempty"`
	Content    string    `json:"content, omitempty"`
	AuthorId   uint64    `json:"authorId, omitempty"`
	AuthorNick string    `json:"AuthorNick, omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"CreatedAt, omitempty"`
}

func (publication *Publication) PreparePublication() error {
	if err := publication.validatePublication(); err != nil {
		return err
	}

	publication.formatePublication()

	return nil
}

func (publication *Publication) validatePublication() error {
	if publication.Title == " " {
		return errors.New(" Title can´t be empty")
	}
	if publication.Content == " " {
		return errors.New(" Content can´t be empty")
	}

	return nil

}

func (publication *Publication) formatePublication() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)

}