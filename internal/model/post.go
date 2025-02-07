package model

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Post struct {
	ID        int                 `json:"id"`
	Info      PostInfo            `json:"-"`
	CreatedAt timestamp.Timestamp `json:"createdAt"` // todo мб надо created_at
	UpdatedAt timestamp.Timestamp `json:"updatedAt"`
}

type PostInfo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  int    `json:"userId"`
	Rating  int    `json:"rating"`
}
