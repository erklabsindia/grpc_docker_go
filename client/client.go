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
	res, err := client.ListPosts(ctx, &post.ListPostsRequest{PageNo: 0, PageSize: 10, Uuid: "", Filter: "", Query: ""})
	if err != nil {
		return errors.Wrap(err, "calling 'client.ListPostsRequest()'")
	}
	fmt.Println(res.Message)
	fmt.Println(res.Status)
	fmt.Println(res.Type)
	for _, pst := range res.Post {
		fmt.Println("----------------------------")
		fmt.Println("User: ", pst.PostedBy.Name)
		fmt.Println("Post: ", pst.Content)
		for _, op := range pst.Options {
			fmt.Println("OptionId: ", op.OptionId)
			fmt.Println("TotalLikes: ", op.TotalLikes)
		}
	}
	return nil
}
