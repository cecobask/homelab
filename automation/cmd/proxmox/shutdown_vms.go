package proxmox

import (
	"context"
	"github.com/cecobask/homelab/automation/cmd"
	"github.com/cecobask/homelab/automation/pkg/proxmox"
	"github.com/spf13/cobra"
	"log/slog"
)

func shutdownVM(ctx context.Context, logger *slog.Logger) *cobra.Command {
	command := &cobra.Command{
		Use:     cmd.CommandNameShutdownVMs,
		Aliases: []string{cmd.CommandAliasShutdownVMs},
		Short:   "Shut down virtual machines",
		RunE: func(c *cobra.Command, args []string) error {
			baseURL, _ := c.Flags().GetString(cmd.FlagNameBaseURL)
			client := proxmox.NewClient(baseURL, logger)
			node, _ := c.Flags().GetString(cmd.FlagNameNode)
			vmids, _ := c.Flags().GetStringSlice(cmd.FlagNameVMIDs)
			return client.ShutdownVMs(ctx, node, vmids)
		},
	}
	command.Flags().String(cmd.FlagNameNode, "", "cluster node name")
	command.Flags().StringSlice(cmd.FlagNameVMIDs, nil, "virtual machine identifiers")
	_ = command.MarkFlagRequired(cmd.FlagNameNode)
	_ = command.MarkFlagRequired(cmd.FlagNameVMIDs)
	return command
}
