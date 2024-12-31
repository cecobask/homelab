package root

import (
	"context"
	"github.com/cecobask/homelab/automation/cmd"
	"github.com/cecobask/homelab/automation/cmd/proxmox"
	"github.com/cecobask/homelab/automation/cmd/tailscale"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
)

func NewCommand(ctx context.Context, logger *slog.Logger) *cobra.Command {
	command := &cobra.Command{
		Use:     cmd.CommandNameRoot,
		Aliases: []string{cmd.CommandAliasRoot},
		Short:   "homelab automation command line interface",
		PersistentPreRun: func(c *cobra.Command, args []string) {
			c.SetOut(os.Stdout)
			c.SetErr(os.Stderr)
		},
		RunE: func(c *cobra.Command, args []string) error {
			return c.Help()
		},
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		SilenceUsage: true,
	}
	command.SetHelpCommand(&cobra.Command{
		Hidden: true,
	})
	command.AddCommand(
		proxmox.NewCommand(ctx, logger),
		tailscale.NewCommand(ctx, logger),
	)
	return command
}
