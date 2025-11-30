package volume

import (
	"fmt"

	"github.com/cecobask/homelab/pkg/client"
	"github.com/spf13/cobra"
)

func newCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		RunE: runCreate,
	}
	cmd.Flags().String("format", "raw", "format of the volume")
	cmd.Flags().String("size", "", "size of the volume (1024M, 1G)")
	cmd.MarkFlagRequired("size")
	return cmd
}

func runCreate(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	node, _ := cmd.Flags().GetString("node")
	size, _ := cmd.Flags().GetString("size")
	storage, _ := cmd.Flags().GetString("storage")
	volume, _ := cmd.Flags().GetString("volume")
	vmid, err := validateVolume(volume)
	if err != nil {
		return err
	}
	params := client.CreateVolumeParams{
		Filename: volume,
		Node:     node,
		Size:     size,
		Storage:  storage,
		VMID:     vmid,
		Format:   format,
	}
	if err = client.New().EnsureVolume(cmd.Context(), params); err != nil {
		return fmt.Errorf("failed to create volume: %w", err)
	}
	return nil
}
