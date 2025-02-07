package convertor

import (
	"post/internal/model"
	"post/pkg/post_v1"
)

func ToPostInfoFromApi(r *post_v1.PostInfoRequest) *model.PostInfo {
	return &model.PostInfo{
		Title:   r.Title,
		Content: r.Content,
		UserId:  int(r.UserId),
	}
}

func ToApiFromPost(m *model.Post) *post_v1.Post {
	return &post_v1.Post{
		Id:        int64(m.ID),
		Info:      ToApiFromPostInfo(m.Info),
		CreatedAt: &m.CreatedAt,
		UpdatedAt: &m.UpdatedAt,
	}
}

func ToApiFromPostInfo(m model.PostInfo) *post_v1.PostInfo {
	return &post_v1.PostInfo{
		Title:   m.Title,
		Content: m.Content,
		UserId:  int64(m.UserId),
		Rating:  int64(m.Rating),
	}
}
