package post

import "post/internal/repository"

type service struct {
	repo repository.PostRepository
}

func NewService(repo repository.PostRepository) repository.PostRepository {
	return &service{repo: repo}
}
