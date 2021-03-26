package cmd

import (
	"github.com/spf13/cobra"
)

// MainCommand performs database migrations
func MainCommand() *cobra.Command {
	var rootCommand = &cobra.Command{
		Use:   "db",
		Short: "Database tooling",
	}

	rootCommand.AddCommand(&cobra.Command{
		Use:   "migrate",
		Short: "Performs database migrations or read migration info",
		Run:   performMigrations,
	})
	return rootCommand
}

func performMigrations(cmd *cobra.Command, args []string) {
	// Here is where you should add the command to perform database migrations
}
