package main

import (
	"context"
	"fmt"
	"log"
	"os"
	post "worklen/proto/proto/posts"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	log := log.New(os.Stdout, "GRPC CLIENT : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	if err := run(ctx, log); err != nil {
		log.Println("main: error:", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, log *log.Logger) error {
	log.Println("main: Initializing GRPC client")
	defer log.Println("main: Completed")
	conn, err := grpc.Dial("localhost:4040", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		errors.Wrap(err, "dialing")
	}
	client := post.NewProtobufServiceClient(conn)
	res, err := client.ListPosts(ctx, &post.ListPostsRequest{PageNo: 1, PageSize: 10, Uuid: "sssss", Filter: "fff", Query: "fff"})
	if err != nil {
		return errors.Wrap(err, "calling 'client.ListPostsRequest()'")
	}
	fmt.Println(res.Message)
	fmt.Println(res.Status)
	fmt.Println(res.Type)
	for i, pst := range res.Post {
		fmt.Println("Index", i)
		fmt.Println("Centent: ", pst.Content)
		fmt.Println("CreatedOn: ", pst.CreatedOn)
	}
	return nil
}
