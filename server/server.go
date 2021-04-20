package server

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"os"
)

func NewGRPCServer(alro *viper.Viper) (g *grpc.Server) {
	g = grpc.NewServer()
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
