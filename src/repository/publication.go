package repository

import (
	"api/src/modells"
	"database/sql"
)

type Publications struct {
	db *sql.DB
}

func NewPublicationRepository(db *sql.DB) *Publications {
	return &Publications{db}
}

func (repository Publications) CreatePublication(publication modells.Publication) (uint64, error) {
	statement, err := repository.db.Prepare("insert into publications(title, content, authorId) values(?,?,?)")
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(publication.Title, publication.Content, publication.AuthorId)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}