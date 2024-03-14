package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net"

	news "worklen/proto/news"

	"worklen/configreader"
	"worklen/news_api"

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
}

type server struct {
	listener   net.Listener
	grpcServer *grpc.Server
	newsApi    news_api.NewsAPI
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
	server.listener = listener
	config, err := configreaderReadEnv()
	if err != nil {
		return server, errors.Wrap(err, "reading env vars")
	}
	server.newsApi = news_api.NewNewsAPI(config.NewsBaseUrl, config.NewsApiKey, config.NewsHttpTimeout)
	server.grpcServer = grpc.NewServer()
	news.RegisterProtobufServiceServer(server.grpcServer, server)
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
