package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	servicepb "github.com/matttproud/contextualizing-contexts/proto"
)

func f(ctx context.Context) {
	log.Println("[BEGIN] f")
	defer log.Println("[END] f")
	// Flimsily make it improbable for this function to continue while the
	// handler is serving.
	time.Sleep(time.Second)
	select {
	case <-time.After(5 * time.Second):
		log.Println("5s")
	case <-ctx.Done():
		log.Println("canceled")
	}
}

type server struct {
	servicepb.UnimplementedTestServer // What kind of viral nonsense of an antipattern is this?
}

func (server) Exercise(ctx context.Context, req *servicepb.Request) (*servicepb.Response, error) {
	log.Println("[BEGIN] Serving gRPC")
	defer log.Println("[DONE] Serving gRPC")
	go f(ctx)
	return new(servicepb.Response), nil
}

func main() {
	l, err := net.Listen("tcp", ":8080")
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