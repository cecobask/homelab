package cmd

import (
	"context"
	"log/slog"

	"github.com/cecobask/homelab/cmd/volume"
	"github.com/spf13/cobra"
)

func newCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "proxmox",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	cmd.AddCommand(volume.NewCommand())
	return cmd
}

func Execute() int {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := newCommand().ExecuteContext(ctx); err != nil {
		slog.Error("failed to execute command", "error", err)
		return 1
	}
	return 0
}
