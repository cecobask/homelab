package proxmox

import (
	"context"
	"github.com/cecobask/homelab/automation/cmd"
	"github.com/cecobask/homelab/automation/pkg/proxmox"
	"github.com/spf13/cobra"
	"log/slog"
)

func deleteVolumes(ctx context.Context, logger *slog.Logger) *cobra.Command {
	command := &cobra.Command{
		Use:     cmd.CommandNameDeleteVolumes,
		Aliases: []string{cmd.CommandAliasDeleteVolumes},
		Short:   "Delete storage volumes",
		RunE: func(c *cobra.Command, args []string) error {
			baseURL, _ := c.Flags().GetString(cmd.FlagNameBaseURL)
			client := proxmox.NewClient(baseURL, logger)
			storageFilters := buildStorageFilters(c.Flags())
			storages, err := client.ListStorages(ctx, storageFilters)
			if err != nil {
				return err
			}
			storageContentFilters := buildStorageContentFilters(c.Flags())
			storageContents, err := client.ListStorageContents(ctx, storages, storageContentFilters)
			if err != nil {
				return err
			}
			return client.DeleteVolumes(ctx, storageContents)
		},
	}
	command.Flags().StringSlice(cmd.FlagNameContentTypes, nil, "content types")
	command.Flags().StringSlice(cmd.FlagNameNodes, nil, "cluster node names")
	command.Flags().StringSlice(cmd.FlagNameStorages, nil, "storage identifiers")
	command.Flags().StringSlice(cmd.FlagNameVolumes, nil, "volume identifiers")
	return command
}
