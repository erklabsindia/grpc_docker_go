package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net"

	apis "worklen/apis"
	news_api "worklen/apis/news"
	post_apis "worklen/apis/posts"
	"worklen/configreader"
	news "worklen/proto/proto/news"
	post "worklen/proto/proto/posts"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	netListen           = net.Listen
	configreaderReadEnv = configreader.ReadEnv
	jsonMarshal         = json.Marshal
	protojsonUnmarshal  = protojson.Unmarshal
)

type Server interface {
	Serve() error
	GracefulStop()
	GetNewsArticles(ctx context.Context, in *news.GetNewsRequest) (*news.NewsList, error)
	ListPosts(ctx context.Context, in *post.ListPostsRequest) (*post.PostResponse, error)
	CreatePost(ctx context.Context, in *post.Post) (*post.PostResponse, error)
	GetPost(ctx context.Context, in *post.GetPostRequest) (*post.PostResponse, error)
	UpdatePost(ctx context.Context, in *post.Post) (*post.PostResponse, error)
	DeletePost(ctx context.Context, in *post.DeletePostRequest) (*post.PostResponse, error)
}

type server struct {
	listener   net.Listener
	grpcServer *grpc.Server
	newsApi    news_api.NewsAPI
	db         *apis.PostgresDb
}

func (s *server) Serve() error {
	return s.grpcServer.Serve(s.listener)
}

func (s *server) GracefulStop() {
	s.grpcServer.GracefulStop()
}

// NewServer creates a new gRPC server.
func NewServer(port int) (Server, error) {
	server := new(server)
	listener, err := netListen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return server, errors.Wrap(err, "tcp listening")
	}
	ctx := context.TODO()
	server.listener = listener
	config, err := configreaderReadEnv()
	if err != nil {
		return server, errors.Wrap(err, "reading env vars")
	}

	server.grpcServer = grpc.NewServer()
	server.newsApi = news_api.NewNewsAPI(config.NewsBaseUrl, config.NewsApiKey, config.NewsHttpTimeout)
	pdb, err := apis.Connect(ctx, config.SQL_HOST, config.SQL_USERNAME, config.SQL_PASS, "work_test", config.SQL_PORT)
	if err != nil {
		return server, errors.Wrap(err, "Error connect database")
	}
	server.db = pdb
	news.RegisterProtobufServiceServer(server.grpcServer, server)
	post.RegisterProtobufServiceServer(server.grpcServer, server)
	reflection.Register(server.grpcServer)
	return server, nil
}

func (s *server) GetNewsArticles(ctx context.Context, in *news.GetNewsRequest) (*news.NewsList, error) {
	pbArticleList := new(news.NewsList)
	articles, err := s.newsApi.GetNews(string(in.Query), int(in.PageSize))
	if err != nil {
		return pbArticleList, errors.Wrap(err, "requesting articles")
	}
	json, err := jsonMarshal(articles)
	if err != nil {
		return pbArticleList, errors.Wrap(err, "marshalling json")
	}
	err = protojsonUnmarshal(json, pbArticleList)
	if err != nil {
		return pbArticleList, errors.Wrap(err, "unmarshalling proto")
	}
	return pbArticleList, nil
}

func (s *server) ListPosts(ctx context.Context, in *post.ListPostsRequest) (*post.PostResponse, error) {
	return post_apis.ListPosts(ctx, s.db, in)
}
func (s *server) CreatePost(ctx context.Context, in *post.Post) (*post.PostResponse, error) {
	return post_apis.CreatePost(ctx, s.db, in)
}
func (s *server) GetPost(ctx context.Context, in *post.GetPostRequest) (*post.PostResponse, error) {
	return post_apis.GetPost(ctx, s.db, in)
}
func (s *server) UpdatePost(ctx context.Context, in *post.Post) (*post.PostResponse, error) {
	return post_apis.UpdatePost(ctx, s.db, in)
}
func (s *server) DeletePost(ctx context.Context, in *post.DeletePostRequest) (*post.PostResponse, error) {
	return post_apis.DeletePost(ctx, s.db, in)
}
