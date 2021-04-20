package main

import (
	"alro/config"
	"alro/server"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"os"
)
var log *logrus.Logger

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	log = logrus.StandardLogger()
	alro := config.Alro
	g := server.NewGRPCServer(alro)

	startServer(alro, g)
}

func startServer(alro *viper.Viper, g *grpc.Server) {
	address := fmt.Sprintf("0.0.0.0:%d", alro.GetInt("port"))
	fmt.Printf("start ALRO!!!")

	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Println("End to bind for gRPC server", "err", err)
		os.Exit(1)
	}
	if err := g.Serve(l); err != nil {
		log.Println("End gRPC server", "err", err)
		os.Exit(1)
	}
}