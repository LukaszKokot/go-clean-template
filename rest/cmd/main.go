package cmd

import (
	"github.com/LukaszKokot/go-clean-template/config"
	"github.com/LukaszKokot/go-clean-template/rest"

	"github.com/spf13/cobra"
)

// MainCommand starts the web server supporting the REST API
func MainCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "web",
		Short: "Starts the Web server",
		Run:   bootWebServer,
	}
}

func bootWebServer(cmd *cobra.Command, args []string) {
	// Here is where we configure our web server and start it
	rest.NewServer(rest.Options{
		Web: config.GetWebServerConfiguration(),
	})
}
