package repository

import "github.com/Harishraju04/Ainyx-Go-Lang-Assignment/db/sqlc"

type Repository struct {
	q *sqlc.Queries
}

func NewRepository(q *sqlc.Queries) *Repository {
	return &Repository{
		q: q,
	}
}
