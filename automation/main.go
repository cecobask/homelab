package main

import (
	"context"
	"github.com/cecobask/homelab/automation/cmd/root"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	err := root.NewCommand(ctx).Execute()
	if err != nil {
		os.Exit(1)
	}
}
