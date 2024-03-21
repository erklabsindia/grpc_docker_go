package server

import (
	"context"
	"fmt"
	"net"

	config_reader "grpc.worklen.com/config"
	gen "grpc.worklen.com/proto/generated"
	database "grpc.worklen.com/server/apis"
	news_api "grpc.worklen.com/server/apis/news"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server interface {
	Serve() error
	GracefulStop()
}

// NewServer creates a new gRPC server.
func StartServer() (Server, error) {
	server := new(ServerInstance)
	ctx := context.TODO()

	config, err := config_reader.ReadEnv()
	if err != nil {
		return server, errors.Wrap(err, "Error: reading env vars")
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 4040))
	if err != nil {
		return server, errors.Wrap(err, "Error: Tcp Listening")
	}
	server.listener = listener

	server.grpcServer = grpc.NewServer()
	server.newsClient = news_api.GetNewsApiClient(config.NewsBaseUrl, config.NewsApiKey, config.NewsHttpTimeout)
	pdb, err := database.Connect(ctx, config.SQL_HOST, config.SQL_USERNAME, config.SQL_PASS, config.SQL_DB, config.SQL_PORT)
	if err != nil {
		return server, errors.Wrap(err, "Error: connect database")
	}
	server.db = pdb
	gen.RegisterProtobufServiceServer(server.grpcServer, server)
	gen.RegisterPostProtobufServiceServer(server.grpcServer, server)
	reflection.Register(server.grpcServer)
	return server, nil
}
