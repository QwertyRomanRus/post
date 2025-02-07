package repository

import (
	"context"
	"post/internal/model"
)

type PostRepository interface {
	Create(context.Context, *model.PostInfo) (*model.Post, error)
	Get(context.Context, int) (*model.Post, error)
}
