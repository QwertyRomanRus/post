package convertor

import (
	"post/internal/model"
	modelRepo "post/internal/repository/post/model"
)

func ToModelFromRepo(post *modelRepo.Post) *model.Post {
	return &model.Post{
		ID:   post.ID,
		Info: ToModelInfoFromRepo(post.Info),
	}
}

func ToModelInfoFromRepo(info modelRepo.PostInfo) model.PostInfo {
	return model.PostInfo{
		Title:   info.Title,
		Content: info.Content,
		UserId:  info.UserId,
		Rating:  info.Rating,
	}
}
