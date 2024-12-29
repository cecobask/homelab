package tailscale

import (
	"context"
	"fmt"
	"github.com/cecobask/homelab/automation/cmd"
	"github.com/cecobask/homelab/automation/pkg/tailscale"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"log/slog"
)

type contextKey string

const contextKeyFilters = "filters"

func NewCommand(ctx context.Context, logger *slog.Logger) *cobra.Command {
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
		deleteDevices(ctx, logger),
		listDevices(ctx, logger),
	)
	command.PersistentFlags().String(cmd.FlagNameBaseURL, cmd.BaseURLTailscale, "base url for the tailscale api")
	return command
}

func buildDeviceFilters(flags *pflag.FlagSet) (*tailscale.DeviceFilters, error) {
	ids, _ := flags.GetStringSlice(cmd.FlagNameIDs)
	hostnames, _ := flags.GetStringSlice(cmd.FlagNameHostnames)
	tags, _ := flags.GetStringSlice(cmd.FlagNameTags)
	return tailscale.NewDeviceFilters(ids, hostnames, tags), nil
}
