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
		Use:     cmd.CommandNameShutdownVM,
		Aliases: []string{cmd.CommandAliasShutdownVM},
		Short:   "Shut down virtual machine",
		RunE: func(c *cobra.Command, args []string) error {
			baseURL, _ := c.Flags().GetString(cmd.FlagNameBaseURL)
			client := proxmox.NewClient(baseURL, logger)
			node, _ := c.Flags().GetString(cmd.FlagNameNode)
			vmid, _ := c.Flags().GetString(cmd.FlagNameVMID)
			return client.ShutdownVM(ctx, node, vmid)
		},
	}
	command.Flags().String(cmd.FlagNameNode, "", "cluster node name")
	command.Flags().String(cmd.FlagNameVMID, "", "virtual machine id")
	_ = command.MarkFlagRequired(cmd.FlagNameNode)
	_ = command.MarkFlagRequired(cmd.FlagNameVMID)
	return command
}
