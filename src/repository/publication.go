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

func (repository Publications) SearchPublicationById(publicationId uint64) (modells.Publication, error) {

	rows, err := repository.db.Query(
		`SELECT  p.*, u.nick, from publications p inner join users u on u.id = p.authorId = ?`, publicationId,
	)
	if err != nil {
		return modells.Publication{}, err
	}

	defer rows.Close()

	var publication modells.Publication
	if rows.Next() {
		if err = rows.Scan(
			&publication.Id,
			&publication.Title,
			&publication.Content,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return modells.Publication{}, err
		}
	}

	return publication, nil

}

func (repository Publications) SearchPublications(userId uint64) ([]modells.Publication, error) {

	rows, err := repository.db.Query(`
		select distinct p.* , u.nick, from publications p 
		inner join users u on u.id=p.authorId
		inner join followers s on p.authorId = s.user_id
		where u.id = ? or s.followerId = ?`,
		userId, userId,
	)

	if err != nil {
		return []modells.Publication{}, err
	}
	var publications []modells.Publication
	for rows.Next() {
		var publication modells.Publication

		if err = rows.Scan(
			&publication.Id,
			&publication.Title,
			&publication.Content,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)

	}

	return publications, nil
}

func (repository Publications) UpdatePublication(publicationId uint64, publication modells.Publication) error {

	statement, err := repository.db.Prepare("update publications set title = ?, content = ? where id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(publication.Title, publication.Content, publicationId); err != nil {
		return err
	}

	return nil

}

func (repository Publications) DeletePublication(publicationId uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM publications WHERE id = ?")
	if err != nil {
		return err
	}

	if _, err = statement.Exec(publicationId); err != nil {
		return err
	}

	return nil
}

func (repository Publications) SearchPublicationsByUser(userId uint64) ([]modells.Publication, error) {

	rows, err := repository.db.Query(
		`
	SELECT p.* , u.nick from publications p 
	JOIN users u ON u.id = p.authorId
	WHERE p.authorId = ?
`,
		userId,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var publications []modells.Publication

	for rows.Next() {
		var publication modells.Publication
		if err = rows.Scan(
			&publication.Id,
			&publication.Title,
			&publication.Content,
			&publication.AuthorId,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil

}
