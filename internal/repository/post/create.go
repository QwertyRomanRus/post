package post

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"post/internal/model"
)

func (r *repo) Create(ctx context.Context, req *model.PostInfo) (*model.Post, error) {
	sql, args, err := sq.Insert(tableName).
		Columns(titleColumn, contentColumn, userIdColumn, ratingColumn).
		PlaceholderFormat(sq.Dollar).
		Values(req.Title, req.Content, req.UserId, defaultRatingValue).
		Suffix("RETURNING " + idColumn).
		ToSql()

	if err != nil {
		return nil, err
	}

	var id int
	err = r.db.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return nil, err
	}

	return r.Get(ctx, id)
}
