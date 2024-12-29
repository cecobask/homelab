package proxmox

import (
	"context"
	"fmt"
	"github.com/cecobask/homelab/automation/cmd"
	"github.com/spf13/cobra"
	"log/slog"
)

func NewCommand(ctx context.Context, logger *slog.Logger) *cobra.Command {
	command := &cobra.Command{
		Use:     fmt.Sprintf("%s [command]", cmd.CommandNameProxmox),
		Aliases: []string{cmd.CommandAliasProxmox},
		Short:   "Proxmox commands",
		RunE: func(c *cobra.Command, args []string) error {
			return c.Help()
		},
		SilenceUsage: true,
	}
	command.AddCommand(
		deleteVolume(ctx, logger),
		destroyVM(ctx, logger),
		shutdownVM(ctx, logger),
		stopVM(ctx, logger),
	)
	command.PersistentFlags().String(cmd.FlagNameBaseURL, cmd.BaseURLProxmox, "base url for the proxmox api")
	return command
}
