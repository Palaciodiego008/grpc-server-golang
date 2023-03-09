package main

import (
	"bookshop/server/bookshop/pb"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedInventoryServer
}

func getSampleBooks() []*pb.Book {
	books := []*pb.Book{
		{
			Title:     "The Hitchhiker's Guide to the Galaxy",
			Author:    "Douglas Adams",
			PageCount: 224,
		},

		{
			Title:     "The Restaurant at the End of the Universe",
			Author:    "Jerry Pournelle",
			PageCount: 392,
		},

		{
			Title:     "Life, the Universe and Everything",
			Author:    "Gabriel García Márquez",
			PageCount: 432,
		},
	}

	return books
}

func (s *Server) GetBookList(ctx context.Context, in *pb.GetBookListRequest) (*pb.GetBookListResponse, error) {
	gblResponse := &pb.GetBookListResponse{Books: getSampleBooks()}
	return gblResponse, nil
}

func main() {

	address := ":50051"

	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	/// check if the server is running and then show and message with the port
	log.Println("Server is running on port " + address)
	fmt.Println("Running gRPC server on port 50051")

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	pb.RegisterInventoryServer(grpcServer, &Server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
