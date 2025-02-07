package post

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"post/internal/model"
	"post/internal/repository/post/convertor"
	modelRepo "post/internal/repository/post/model"
)

func (r *repo) Get(ctx context.Context, id int) (*model.Post, error) {
	sql, args, err := sq.Select(idColumn, titleColumn, contentColumn, userIdColumn, ratingColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		Limit(1).
		ToSql()

	if err != nil {
		return nil, err
	}

	var postRepo modelRepo.Post
	err = r.db.QueryRow(ctx, sql, args...).
		Scan(&postRepo.ID, &postRepo.Info.Title, &postRepo.Info.Content, &postRepo.Info.UserId, &postRepo.Info.Rating)

	if err != nil {
		return nil, err
	}

	return convertor.ToModelFromRepo(&postRepo), nil
}
