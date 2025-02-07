package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"post/internal/api/post"
	"post/internal/config"
	postRepo "post/internal/repository/post"
	postSvc "post/internal/service/post"
	"post/pkg/post_v1"
)

const (
	EnvPath = ".env"
)

func main() {
	err := config.Load(EnvPath)
	ctx := context.Background()
	if err != nil {
		log.Fatal("Error loading .env: ", err)
	}
	pgConf, err := config.NewPGConfig()
	if err != nil {
		log.Fatal("Error loading pg config: ", err)
	}

	pgCon, err := pgxpool.Connect(ctx, pgConf.GetDsn())
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer pgCon.Close()

	err = pgCon.Ping(ctx)
	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Error grpc listening on port: ", err)
	}

	defer l.Close()
	// init user
	postRepository := postRepo.NewRepository(pgCon)
	postService := postSvc.NewService(postRepository)
	postApi := post.Implementation{PostService: postService}

	s := grpc.NewServer()
	post_v1.RegisterPostServiceV1Server(s, &postApi)

	reflection.Register(s)

	err = s.Serve(l)
	if err != nil {
		log.Fatal("Error grpc listening on port: ", err)
	}
}
