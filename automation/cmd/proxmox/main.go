package proxmox

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/cecobask/homelab/automation/cmd"
	"github.com/cecobask/homelab/automation/pkg/proxmox"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewCommand(ctx context.Context, logger *slog.Logger) *cobra.Command {
	command := &cobra.Command{
		Use:     fmt.Sprintf("%s [command]", cmd.CommandNameProxmox),
		Aliases: []string{cmd.CommandAliasProxmox},
		Short:   "Proxmox commands",
		RunE: func(c *cobra.Command, args []string) error {
			return c.Help()
		},
		SilenceUsage: true,
	}
	command.AddCommand(
		deleteVolumes(ctx, logger),
		destroyVMs(ctx, logger),
		listVMs(ctx, logger),
		stopVMs(ctx, logger),
	)
	command.PersistentFlags().String(cmd.FlagNameBaseURL, cmd.ProxmoxBaseURL, "base url for the proxmox api")
	return command
}

func buildVirtualMachineFilters(flags *pflag.FlagSet) proxmox.VirtualMachineFilters {
	ids, _ := flags.GetIntSlice(cmd.FlagNameIDs)
	nodes, _ := flags.GetStringSlice(cmd.FlagNameNodes)
	tags, _ := flags.GetStringSlice(cmd.FlagNameTags)
	return proxmox.NewVirtualMachineFilters(ids, nodes, tags)
}

func buildStorageFilters(flags *pflag.FlagSet) proxmox.StorageFilters {
	nodes, _ := flags.GetStringSlice(cmd.FlagNameNodes)
	storages, _ := flags.GetStringSlice(cmd.FlagNameStorages)
	return proxmox.NewStorageFilters(nodes, storages)
}

func buildStorageContentFilters(flags *pflag.FlagSet) proxmox.StorageContentFilters {
	contentTypes, _ := flags.GetStringSlice(cmd.FlagNameContentTypes)
	volumes, _ := flags.GetStringSlice(cmd.FlagNameVolumes)
	return proxmox.NewStorageContentFilters(contentTypes, volumes)
}
