package tailscale

import (
	"context"
	"fmt"
	"github.com/cecobask/homelab/automation/cmd"
	"github.com/cecobask/homelab/automation/pkg/tailscale"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type contextKey string

const filtersContextKey = "filters"

func NewCommand(ctx context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:     fmt.Sprintf("%s [command]", cmd.CommandNameTailscale),
		Aliases: []string{cmd.CommandAliasTailscale},
		Short:   "Tailscale commands",
		RunE: func(c *cobra.Command, args []string) error {
			return c.Help()
		},
		SilenceUsage: true,
	}
	command.AddCommand(
		deleteDevices(ctx),
		listDevices(ctx),
	)
	return command
}

func buildDeviceFilters(flags *pflag.FlagSet) (*tailscale.DeviceFilters, error) {
	ids, err := flags.GetStringSlice(cmd.FlagNameIDs)
	if err != nil {
		return nil, err
	}
	hostnames, err := flags.GetStringSlice(cmd.FlagNameHostnames)
	if err != nil {
		return nil, err
	}
	tags, err := flags.GetStringSlice(cmd.FlagNameTags)
	if err != nil {
		return nil, err
	}
	return tailscale.NewDeviceFilters(ids, hostnames, tags), nil
}
