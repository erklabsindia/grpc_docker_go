package main

import (
	"context"
	"fmt"
	"log"
	"os"
	news "worklen/proto/news"

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
	client := news.NewProtobufServiceClient(conn)
	res, err := client.GetNewsArticles(ctx, &news.GetNewsRequest{Query: string("bitcoin"), PageSize: int32(2)})
	if err != nil {
		return errors.Wrap(err, "calling 'client.GetNewsRequest()'")
	}
	for _, article := range res.Articles {
		fmt.Println("\nTitle: ", article.Title)
		fmt.Println("Author: ", article.Author)
		fmt.Print("\n")
	}
	return nil
}
