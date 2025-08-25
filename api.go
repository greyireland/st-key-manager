package main

import (
	"flag"
	"fmt"
	"github.com/greyireland/log"
	"google.golang.org/grpc"
	"net"
	"st-key-manager/internal/config"
	"st-key-manager/internal/server"
	"st-key-manager/internal/svc"
	"st-key-manager/keymanager"
	"st-key-manager/pkg/conf"
	"st-key-manager/pkg/jsonx"
)

var configFile = flag.String("f", "etc/st-key-manager.yaml", "the config file")

func main() {
	flag.Parse()
	log.Root().SetHandler(log.DefaultFileHandler())

	//log.Root().SetHandler(log.StdoutHandler)
	var c config.Config
	conf.MustLoad(&c, *configFile)
	log.Info("config", "conf", jsonx.JSON(c))
	//metrics.Setup(c.MetricsOn)
	ctx := svc.NewServiceContext(c)
	InitServer(&c, ctx)
}
func InitServer(c *config.Config, ctx *svc.ServiceContext) {
	lis, err := net.Listen("tcp", c.ListenOn)
	if err != nil {
		panic(err)
	}
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(svc.EnsureValidToken),
		grpc.StreamInterceptor(svc.EnsureValidTokenStream),
	}
	grpcServer := grpc.NewServer(opts...)
	serverImpl := server.NewKeymanagerServer(ctx)
	keymanager.RegisterKeymanagerServer(grpcServer, serverImpl)
	fmt.Printf("Start at %s \n", c.ListenOn)
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Println("Serve error", err)
	}
}
