package main

import (
	"alro/config"
	"alro/server"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	//log := logrus.StandardLogger()
	alro := config.Alro
	server.NewGRPCServer(alro)
}
