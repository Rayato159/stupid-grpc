package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Rayato159/stupid-grpc/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTransferServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:443")
	if err != nil {
		log.Fatalf("error, failed to listen: %v", err)
	}
	log.Printf("success, server is starting on 0.0.0.0:443")

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterTransferServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}

func (s *server) FindOneProduct(ctx context.Context, in *pb.ProductReq) (*pb.ProductRes, error) {
	if in == nil {
		return nil, fmt.Errorf("nil request reference")
	}

	products := buildProducts()

	if products[int(in.Id)] != nil {
		return &pb.ProductRes{
			Id:    int32(products[int(in.Id)].Id),
			Title: products[int(in.Id)].Title,
		}, nil
	}
	return nil, fmt.Errorf("product_id: %d not found", in.Id)
}

type Product struct {
	Id    int
	Title string
}

func buildProducts() map[int]*Product {
	return map[int]*Product{
		1: {
			Id:    1,
			Title: "Coffee",
		},
		2: {
			Id:    2,
			Title: "Tea",
		},
		3: {
			Id:    3,
			Title: "Coke",
		},
	}
}
