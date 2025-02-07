package post

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"post/internal/repository"
)

const (
	tableName = "posts"

	idColumn        = "id"
	titleColumn     = "title"
	contentColumn   = "content"
	userIdColumn    = "user_id"
	ratingColumn    = "rating"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"

	limit              = 10
	defaultRatingValue = 0
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.PostRepository {
	return &repo{db: db}
}
