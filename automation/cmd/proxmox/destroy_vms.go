package proxmox

import (
	"context"
	"github.com/cecobask/homelab/automation/cmd"
	"github.com/cecobask/homelab/automation/pkg/proxmox"
	"github.com/spf13/cobra"
	"log/slog"
)

func destroyVMs(ctx context.Context, logger *slog.Logger) *cobra.Command {
	command := &cobra.Command{
		Use:     cmd.CommandNameDestroyVMs,
		Aliases: []string{cmd.CommandAliasDestroyVMs},
		Short:   "Destroy virtual machines",
		RunE: func(c *cobra.Command, args []string) error {
			baseURL, _ := c.Flags().GetString(cmd.FlagNameBaseURL)
			client := proxmox.NewClient(baseURL, logger)
			filters := buildVirtualMachineFilters(c.Flags())
			vms, err := client.ListVMs(ctx, filters)
			if err != nil {
				return err
			}
			destroyUnreferencedDisks, _ := c.Flags().GetBool(cmd.FlagNameDestroyUnreferencedDisks)
			purge, _ := c.Flags().GetBool(cmd.FlagNamePurge)
			return client.DestroyVMs(ctx, vms, destroyUnreferencedDisks, purge)
		},
	}
	command.Flags().IntSlice(cmd.FlagNameIDs, nil, "virtual machine identifiers")
	command.Flags().StringSlice(cmd.FlagNameNodes, nil, "cluster node names")
	command.Flags().StringSlice(cmd.FlagNameTags, nil, "virtual machine tags")
	command.Flags().Bool(cmd.FlagNameDestroyUnreferencedDisks, false, "additionally destroy all disks not referenced in the config but with a matching vmid from all enabled storages")
	command.Flags().Bool(cmd.FlagNamePurge, false, "remove vmid from configurations, like backup & replication jobs")
	return command
}
