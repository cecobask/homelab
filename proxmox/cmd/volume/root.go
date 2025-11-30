package volume

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "volume",
	}
	cmd.AddCommand(newCreateCommand())
	cmd.AddCommand(newDeleteCommand())
	cmd.PersistentFlags().String("node", "", "cluster node name")
	cmd.PersistentFlags().String("storage", "", "storage identifier")
	cmd.PersistentFlags().String("volume", "", "volume identifier (vm-<vmid>-<name>)")
	cmd.MarkPersistentFlagRequired("node")
	cmd.MarkPersistentFlagRequired("storage")
	cmd.MarkPersistentFlagRequired("volume")
	return cmd
}

func validateVolume(volume string) (int, error) {
	parts := strings.Split(volume, "-")
	if len(parts) != 3 {
		return 0, errors.New("invalid volume format, expected format is vm-<vmid>-<name>")
	}
	vmid, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid vmid: %s", parts[1])
	}
	return vmid, nil
}
