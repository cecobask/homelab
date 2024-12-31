package tailscale

import (
	"context"
	"fmt"
	"github.com/cecobask/homelab/automation/cmd"
	"github.com/cecobask/homelab/automation/pkg/tailscale"
	"github.com/spf13/cobra"
	"log/slog"
)

func listDevices(ctx context.Context, logger *slog.Logger) *cobra.Command {
	command := &cobra.Command{
		Use:     cmd.CommandNameListDevices,
		Aliases: []string{cmd.CommandAliasListDevices},
		Short:   "List devices w/ optional filters",
		PreRunE: func(c *cobra.Command, args []string) error {
			filters, err := buildDeviceFilters(c.Flags())
			if err != nil {
				return err
			}
			c.SetContext(context.WithValue(c.Context(), contextKey(contextKeyFilters), filters))
			return nil
		},
		RunE: func(c *cobra.Command, args []string) error {
			baseURL, _ := c.Flags().GetString(cmd.FlagNameBaseURL)
			client := tailscale.NewClient(baseURL, logger)
			tailnetName, _ := c.Flags().GetString(cmd.FlagNameTailnetName)
			filters, ok := c.Context().Value(contextKey(contextKeyFilters)).(*tailscale.DeviceFilters)
			if !ok {
				return fmt.Errorf("no device filters found in context")
			}
			_, err := client.ListDevices(ctx, tailnetName, *filters)
			if err != nil {
				return err
			}
			return nil
		},
	}
	command.Flags().String(cmd.FlagNameTailnetName, "-", "tailnet name; '-' dash will reference the default tailnet")
	command.Flags().StringSlice(cmd.FlagNameIDs, nil, "device identifiers")
	command.Flags().StringSlice(cmd.FlagNameHostnames, nil, "device hostnames")
	command.Flags().StringSlice(cmd.FlagNameTags, nil, "device tags")
	return command
}
