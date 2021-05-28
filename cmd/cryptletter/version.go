package cryptletter

import (
	"fmt"
	"github.com/spf13/cobra"
)

// NewCmd builds a new analyse url command
func NewVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "executables version",
		Run:   func(cmd *cobra.Command, args []string) {
			fmt.Println("cryptletter version 3.1.0")
		},
	}

	return cmd
}
