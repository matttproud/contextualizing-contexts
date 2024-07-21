package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	servicepb "github.com/matttproud/contextualizing-contexts/proto"
)

type server struct {
	servicepb.UnimplementedTestServer // What kind of viral nonsense of an antipattern is this?
}

func g() {
	log.Println("g(): context has been canceled")
}

func (server) Exercise(ctx context.Context, req *servicepb.Request) (*servicepb.Response, error) {
	log.Println("[BEGIN] Serving gRPC")
	defer log.Println("[DONE] Serving gRPC")
	context.AfterFunc(ctx, g)
	return new(servicepb.Response), nil
}

func main() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()
	srv := grpc.NewServer()
	servicepb.RegisterTestServer(srv, new(server))
	reflection.Register(srv)
	if err := srv.Serve(l); err != nil {
		log.Fatalln(err)
	}
}
