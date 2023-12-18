package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Gafforov-Bahrom/uzum_admin/internal/config"
	desc "github.com/Gafforov-Bahrom/uzum_admin/pkg/grpc/admin"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

var cfgSrcEnv *string

func init() {
	cfgSrcEnv = flag.String("env", "./.env", "ENV config source")
	flag.Parse()
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.NewConfig(*cfgSrcEnv)
	if err != nil {
		panic(err)
	}

	sp := NewServiceProvider(cfg)

	server := grpc.NewServer()
	server.RegisterService(&desc.AdminService_ServiceDesc, sp.GetProductServer())

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		panic(err)
	}

	go func() {
		<-ctx.Done()
		fmt.Println("context cancelled")
		server.GracefulStop()
	}()

	httpServer := sp.GetHTTPServer()

	wg.Add(1)
	go httpServer.Start(ctx, &wg)
	server.Serve(lis)
	wg.Wait()
}
