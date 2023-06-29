package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Rayato159/stupid-grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	url := "0.0.0.0:443"
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error, failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTransferClient(conn)

	res, err := client.FindOneProduct(ctx, &pb.ProductReq{
		Id: 1,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
