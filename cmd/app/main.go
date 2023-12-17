package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Gafforov-Bahrom/uzum_admin/internal/config"
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

	_ = NewServiceProvider(cfg)
}
