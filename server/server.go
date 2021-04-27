package server

import (
	model "alro/model"
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"time"
)

type EchoServer struct {
	model.EchoClient
}

func NewEchoServer()(*EchoServer, error) {
	return &EchoServer{}, nil
}

func (e *EchoServer) Echo(ctx context.Context, req *model.EchoRequest) (*model.EchoResponse, error) {
	return nil, nil
}

func NewGRPCServer() (*grpc.Server, error) {
	g := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(
				grpc_ctxtags.WithFieldExtractor(
					grpc_ctxtags.CodeGenRequestFieldExtractor,
				),
			),
		),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge: 10 * time.Second,
		}),
	)
	e, err := NewEchoServer()
	if err != nil {
		return nil, nil
	}

	model.RegisterEchoServer(g, e)
	return g, nil
}
