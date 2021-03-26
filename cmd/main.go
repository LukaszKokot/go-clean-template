package cmd

import (
	dbCmd "github.com/LukaszKokot/go-clean-template/db/cmd"
	grpcCmd "github.com/LukaszKokot/go-clean-template/grpc/cmd"
	restCmd "github.com/LukaszKokot/go-clean-template/rest/cmd"

	"github.com/spf13/cobra"
)

var (
	cfgFile         string
	restCommandFunc = restCmd.MainCommand
	dbCommandFunc   = dbCmd.MainCommand
	grpcCommandFunc = grpcCmd.MainCommand
)

// RootCommand returns the main command, properly configured
func RootCommand() *cobra.Command {
	var rootCommand = &cobra.Command{
		Use:              "App",
		PersistentPreRun: initConfig,
	}

	rootCommand.PersistentFlags().StringVar(&cfgFile, "config", "", "/path/to/config.yml")
	rootCommand.AddCommand(restCommandFunc())
	rootCommand.AddCommand(dbCommandFunc())
	rootCommand.AddCommand(grpcCommandFunc())
	return rootCommand
}

func initConfig(ccmd *cobra.Command, args []string) {
	// Read the configuration file(s) here and anything else that
	// is related to general configuration and may be found in the command line
	// args.
}
