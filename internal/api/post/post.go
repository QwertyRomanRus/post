package post

import (
	"post/internal/service"
	"post/pkg/post_v1"
)

type Implementation struct {
	PostService service.PostService
	post_v1.UnimplementedPostServiceV1Server
}
