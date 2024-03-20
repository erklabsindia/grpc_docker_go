package server

import (
	"context"
	"encoding/json"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	gen "grpc.worklen.com/proto/generated"
	database "grpc.worklen.com/server/apis"
	news_api "grpc.worklen.com/server/apis/news"
	post_apis "grpc.worklen.com/server/apis/posts"
)

type ServerInstance struct {
	listener   net.Listener
	grpcServer *grpc.Server
	newsClient news_api.NewsAPI
	db         *database.PostgresDb
}

func (s *ServerInstance) Serve() error {
	return s.grpcServer.Serve(s.listener)
}

func (s *ServerInstance) GracefulStop() {
	s.grpcServer.GracefulStop()
}

func (s *ServerInstance) GetNewsArticles(ctx context.Context, in *gen.GetNewsRequest) (*gen.NewsList, error) {
	pbArticleList := new(gen.NewsList)
	articles, err := s.newsClient.GetNews(string(in.Query), int(in.PageSize))
	if err != nil {
		return pbArticleList, errors.Wrap(err, "requesting articles")
	}
	json, err := json.Marshal(articles)
	if err != nil {
		return pbArticleList, errors.Wrap(err, "marshalling json")
	}
	err = protojson.Unmarshal(json, pbArticleList)
	if err != nil {
		return pbArticleList, errors.Wrap(err, "unmarshalling proto")
	}
	return pbArticleList, nil
}

func (s *ServerInstance) ListPost(ctx context.Context, in *gen.ListPostsRequest) (*gen.PostResponse, error) {
	return post_apis.ListPosts(ctx, s.db, in)
}
func (s *ServerInstance) CreatePost(ctx context.Context, in *gen.Post) (*gen.PostResponse, error) {
	return post_apis.CreatePost(ctx, s.db, in)
}
func (s *ServerInstance) GetPost(ctx context.Context, in *gen.GetPostRequest) (*gen.PostResponse, error) {
	return post_apis.GetPost(ctx, s.db, in)
}
func (s *ServerInstance) UpdatePost(ctx context.Context, in *gen.Post) (*gen.PostResponse, error) {
	return post_apis.UpdatePost(ctx, s.db, in)
}
func (s *ServerInstance) DeletePost(ctx context.Context, in *gen.DeletePostRequest) (*gen.PostResponse, error) {
	return post_apis.DeletePost(ctx, s.db, in)
}
