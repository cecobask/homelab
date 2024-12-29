package proxmox

import (
	"context"
	"github.com/cecobask/homelab/automation/cmd"
	"github.com/cecobask/homelab/automation/pkg/proxmox"
	"github.com/spf13/cobra"
	"log/slog"
)

func deleteVolume(ctx context.Context, logger *slog.Logger) *cobra.Command {
	command := &cobra.Command{
		Use:     cmd.CommandNameDeleteVolume,
		Aliases: []string{cmd.CommandAliasDeleteVolume},
		Short:   "Delete storage volume",
		RunE: func(c *cobra.Command, args []string) error {
			baseURL, _ := c.Flags().GetString(cmd.FlagNameBaseURL)
			client := proxmox.NewClient(baseURL, logger)
			node, _ := c.Flags().GetString(cmd.FlagNameNode)
			storage, _ := c.Flags().GetString(cmd.FlagNameStorage)
			volume, _ := c.Flags().GetString(cmd.FlagNameVolume)
			return client.DeleteVolume(ctx, node, storage, volume)
		},
	}
	command.Flags().String(cmd.FlagNameNode, "", "cluster node name")
	command.Flags().String(cmd.FlagNameStorage, "", "storage identifier")
	command.Flags().String(cmd.FlagNameVolume, "", "volume identifier")
	_ = command.MarkFlagRequired(cmd.FlagNameNode)
	_ = command.MarkFlagRequired(cmd.FlagNameStorage)
	_ = command.MarkFlagRequired(cmd.FlagNameVolume)
	return command
}
