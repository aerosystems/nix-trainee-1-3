package resource

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func (r *Repository) Create(comment Comment) {
	q := "INSERT IGNORE INTO `comments` (id, post_id, name, email, body) VALUES (?, ?, ?, ?, ?);"
	insert, err := r.db.Prepare(q)
	if err != nil {
		panic(err)
	}

	insert.Exec(comment.Id, comment.PostId, comment.Name, comment.Email, comment.Body)
	insert.Close()
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
