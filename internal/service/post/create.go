package post

import (
	"context"
	"post/internal/model"
)

func (s *service) Create(ctx context.Context, info *model.PostInfo) (*model.Post, error) {
	user, err := s.repo.Create(ctx, info)
	if err != nil {
		return nil, err
	}

	return user, nil
}
