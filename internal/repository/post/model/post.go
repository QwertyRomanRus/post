package model

import "github.com/golang/protobuf/ptypes/timestamp"

type Post struct {
	ID        int                  `db:"id"`
	Info      PostInfo             `db:"-"`
	CreatedAt *timestamp.Timestamp `db:"created_at"`
	UpdatedAt *timestamp.Timestamp `db:"updated_at"`
}

type PostInfo struct {
	Title   string `db:"title"`
	Content string `db:"content"`
	UserId  int    `db:"user_id"`
	Rating  int    `db:"rating"`
}
