package proxmox

import (
	"context"
	"github.com/cecobask/homelab/automation/cmd"
	"github.com/cecobask/homelab/automation/pkg/proxmox"
	"github.com/spf13/cobra"
	"log/slog"
)

func listVMs(ctx context.Context, logger *slog.Logger) *cobra.Command {
	command := &cobra.Command{
		Use:     cmd.CommandNameListVMs,
		Aliases: []string{cmd.CommandAliasListVMs},
		Short:   "List virtual machines",
		RunE: func(c *cobra.Command, args []string) error {
			baseURL, _ := c.Flags().GetString(cmd.FlagNameBaseURL)
			client := proxmox.NewClient(baseURL, logger)
			filters := buildVirtualMachineFilters(c.Flags())
			vms, err := client.ListVMs(ctx, filters)
			if err != nil {
				return err
			}
			logger.Info("virtual machines", slog.Any("results", vms))
			return nil
		},
	}
	command.Flags().IntSlice(cmd.FlagNameIDs, nil, "virtual machine identifiers")
	command.Flags().StringSlice(cmd.FlagNameNodes, nil, "cluster node names")
	command.Flags().StringSlice(cmd.FlagNameTags, nil, "virtual machine tags")
	return command
}
