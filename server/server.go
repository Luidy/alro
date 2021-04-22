package server

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"time"
)

func NewGRPCServer() (g *grpc.Server) {
	g = grpc.NewServer(
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

	return g
}
