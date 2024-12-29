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
		Use:     cmd.CommandNameDestroyVM,
		Aliases: []string{cmd.CommandAliasDestroyVM},
		Short:   "Destroy virtual machine",
		RunE: func(c *cobra.Command, args []string) error {
			baseURL, _ := c.Flags().GetString(cmd.FlagNameBaseURL)
			client := proxmox.NewClient(baseURL, logger)
			node, _ := c.Flags().GetString(cmd.FlagNameNode)
			vmid, _ := c.Flags().GetString(cmd.FlagNameVMID)
			destroyUnreferencedDisks, _ := c.Flags().GetBool(cmd.FlagNameDestroyUnreferencedDisks)
			purge, _ := c.Flags().GetBool(cmd.FlagNamePurge)
			return client.DestroyVM(ctx, node, vmid, destroyUnreferencedDisks, purge)
		},
	}
	command.Flags().String(cmd.FlagNameNode, "", "cluster node name")
	command.Flags().String(cmd.FlagNameVMID, "", "virtual machine id")
	command.Flags().Bool(cmd.FlagNameDestroyUnreferencedDisks, false, "additionally destroy all disks not referenced in the config but with a matching vmid from all enabled storages")
	command.Flags().Bool(cmd.FlagNamePurge, false, "remove vmid from configurations, like backup & replication jobs and HA.")
	_ = command.MarkFlagRequired(cmd.FlagNameNode)
	_ = command.MarkFlagRequired(cmd.FlagNameVMID)
	return command
}
