package volume

import (
	"fmt"

	"github.com/cecobask/homelab/pkg/client"
	"github.com/spf13/cobra"
)

func newDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete",
		RunE: runDelete,
	}
	return cmd
}

func runDelete(cmd *cobra.Command, args []string) error {
	node, _ := cmd.Flags().GetString("node")
	storage, _ := cmd.Flags().GetString("storage")
	volume, _ := cmd.Flags().GetString("volume")
	if _, err := validateVolume(volume); err != nil {
		return err
	}
	params := client.DeleteVolumeParams{
		Node:    node,
		Storage: storage,
		Volume:  volume,
	}
	if err := client.New().DeleteVolume(cmd.Context(), params); err != nil {
		return fmt.Errorf("failed to delete volume: %w", err)
	}
	return nil
}
