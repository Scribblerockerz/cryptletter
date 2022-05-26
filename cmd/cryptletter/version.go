package cryptletter

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Version string

// NewVersionCmd lists the current app version
func NewVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "executables version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("cryptletter version %s\n", Version)
		},
	}

	return cmd
}
