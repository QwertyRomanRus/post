package service

import (
	"context"
	"post/internal/model"
)

type PostService interface {
	Create(context.Context, *model.PostInfo) (*model.Post, error)
	//Get(context.Context, int) (*model.Post, error)
}
