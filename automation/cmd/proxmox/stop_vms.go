package proxmox

import (
	"context"
	"github.com/cecobask/homelab/automation/cmd"
	"github.com/cecobask/homelab/automation/pkg/proxmox"
	"github.com/spf13/cobra"
	"log/slog"
)

func stopVMs(ctx context.Context, logger *slog.Logger) *cobra.Command {
	command := &cobra.Command{
		Use:     cmd.CommandNameStopVMs,
		Aliases: []string{cmd.CommandAliasStopVMs},
		Short:   "Stop virtual machines",
		RunE: func(c *cobra.Command, args []string) error {
			baseURL, _ := c.Flags().GetString(cmd.FlagNameBaseURL)
			client := proxmox.NewClient(baseURL, logger)
			filters := buildVirtualMachineFilters(c.Flags())
			vms, err := client.ListVMs(ctx, filters)
			if err != nil {
				return err
			}
			return client.StopVMs(ctx, vms)
		},
	}
	command.Flags().IntSlice(cmd.FlagNameIDs, nil, "virtual machine identifiers")
	command.Flags().StringSlice(cmd.FlagNameNodes, nil, "cluster node names")
	command.Flags().StringSlice(cmd.FlagNameTags, nil, "virtual machine tags")
	return command
}
