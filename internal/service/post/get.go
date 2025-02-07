package post

import (
	"context"
	"post/internal/model"
)

func (s *service) Get(ctx context.Context, id int) (*model.Post, error) {
	post, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return post, nil
}
