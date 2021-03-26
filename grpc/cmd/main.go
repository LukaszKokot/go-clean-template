package cmd

import (
	"github.com/spf13/cobra"
)

// MainCommand starts the GRPC server
func MainCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "web",
		Short: "Starts the GRPC server",
		Run:   bootServer,
	}
}

func bootServer(cmd *cobra.Command, args []string) {
	// Here is where we configure our GRPC server and start it
}
