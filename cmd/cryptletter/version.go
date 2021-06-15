package cryptletter

import (
	"fmt"
	"github.com/spf13/cobra"
)

// NewVersionCmd lists the current app version
func NewVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "executables version",
		Run:   func(cmd *cobra.Command, args []string) {
			fmt.Println("cryptletter version 3.1.1")
		},
	}

	return cmd
}
