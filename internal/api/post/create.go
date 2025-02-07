package post

import (
	"context"
	"post/internal/convertor"
	"post/pkg/post_v1"
)

func (i *Implementation) Create(ctx context.Context, req *post_v1.CreateRequest) (*post_v1.CreateResponse, error) {
	post, err := i.PostService.Create(ctx, convertor.ToPostInfoFromApi(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &post_v1.CreateResponse{Post: convertor.ToApiFromPost(post)}, nil
}
