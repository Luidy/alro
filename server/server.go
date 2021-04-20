package server

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"os"
	"time"
)

func NewGRPCServer(alro *viper.Viper) (g *grpc.Server) {
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
	address := fmt.Sprintf("0.0.0.0:%d", alro.GetInt("port"))
	fmt.Printf(address)

	l, err := net.Listen("tcp", address)
	if err != nil {
		log.WithError(err).Error("End to bind for gRPC server")
		os.Exit(1)
	}
	if err := g.Serve(l); err != nil {
		log.WithError(err).Error("End gRPC server")
		os.Exit(1)
	}

	return g
}
