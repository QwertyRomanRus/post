package post

import (
	"context"
	"post/internal/convertor"
	"post/pkg/post_v1"
)

func (i *Implementation) Get(ctx context.Context, req *post_v1.GetRequest) (*post_v1.GetResponse, error) {
	post, err := i.PostService.Get(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}

	return &post_v1.GetResponse{Post: convertor.ToApiFromPost(post)}, nil
}
