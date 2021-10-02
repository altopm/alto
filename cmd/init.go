package cmd

import (
	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use:                    "init",
	Aliases:                []string{"now"},
	Short:                  "Initialize a new project",
	Long:                   `Initialize a new project to be published to the alto registry, using "sudo alto publish -now"`,
	Example:                "sudo alto init altopm",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

func init() {
	rootCmd.AddCommand(initCommand)
}
