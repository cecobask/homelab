package main

import (
	"context"
	"github.com/cecobask/homelab/automation/cmd/root"
	"github.com/cecobask/homelab/automation/pkg/logger"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	log := logger.NewLogger(os.Stdout)
	err := root.NewCommand(ctx, log).Execute()
	if err != nil {
		os.Exit(1)
	}
}
