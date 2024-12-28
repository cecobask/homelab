package tailscale

import (
	"context"
	"fmt"
	"github.com/cecobask/homelab/automation/cmd"
	"github.com/cecobask/homelab/automation/pkg/tailscale"
	"github.com/spf13/cobra"
)

func listDevices(ctx context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:     fmt.Sprintf("%s [command]", cmd.CommandNameListDevices),
		Aliases: []string{cmd.CommandAliasListDevices},
		Short:   "List devices w/ optional filters",
		PreRunE: func(c *cobra.Command, args []string) error {
			filters, err := buildDeviceFilters(c.Flags())
			if err != nil {
				return err
			}
			c.SetContext(context.WithValue(c.Context(), contextKey(filtersContextKey), filters))
			return nil
		},
		RunE: func(c *cobra.Command, args []string) error {
			client := tailscale.NewClient()
			filters, ok := c.Context().Value(contextKey(filtersContextKey)).(*tailscale.DeviceFilters)
			if !ok {
				return fmt.Errorf("no device filters found in context")
			}
			_, err := client.GetDevices(ctx, *filters)
			if err != nil {
				return err
			}
			return nil
		},
	}
	command.Flags().StringSlice(cmd.FlagNameIDs, nil, "device ids filter")
	command.Flags().StringSlice(cmd.FlagNameHostnames, nil, "device hostnames filter")
	command.Flags().StringSlice(cmd.FlagNameTags, nil, "device tags filter")
	return command
}
