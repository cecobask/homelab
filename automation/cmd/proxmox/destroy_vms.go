package proxmox

import (
	"context"
	"github.com/cecobask/homelab/automation/cmd"
	"github.com/cecobask/homelab/automation/pkg/proxmox"
	"github.com/spf13/cobra"
	"log/slog"
)

func destroyVM(ctx context.Context, logger *slog.Logger) *cobra.Command {
	command := &cobra.Command{
		Use:     cmd.CommandNameDestroyVMs,
		Aliases: []string{cmd.CommandAliasDestroyVMs},
		Short:   "Destroy virtual machines",
		RunE: func(c *cobra.Command, args []string) error {
			baseURL, _ := c.Flags().GetString(cmd.FlagNameBaseURL)
			client := proxmox.NewClient(baseURL, logger)
			node, _ := c.Flags().GetString(cmd.FlagNameNode)
			vmids, _ := c.Flags().GetStringSlice(cmd.FlagNameVMIDs)
			destroyUnreferencedDisks, _ := c.Flags().GetBool(cmd.FlagNameDestroyUnreferencedDisks)
			purge, _ := c.Flags().GetBool(cmd.FlagNamePurge)
			return client.DestroyVMs(ctx, node, vmids, destroyUnreferencedDisks, purge)
		},
	}
	command.Flags().String(cmd.FlagNameNode, "", "cluster node name")
	command.Flags().StringSlice(cmd.FlagNameVMIDs, nil, "virtual machine identifiers")
	command.Flags().Bool(cmd.FlagNameDestroyUnreferencedDisks, false, "additionally destroy all disks not referenced in the config but with a matching vmid from all enabled storages")
	command.Flags().Bool(cmd.FlagNamePurge, false, "remove vmid from configurations, like backup & replication jobs")
	_ = command.MarkFlagRequired(cmd.FlagNameNode)
	_ = command.MarkFlagRequired(cmd.FlagNameVMIDs)
	return command
}
