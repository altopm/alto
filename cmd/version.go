package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Alto and its server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Alto CLI v0.1 -- HEAD")
		fmt.Println("Registry:" /* http.Get("https://registry.altopkg.com/status").Status */)
	},
}

func init() {
	rootCmd.AddCommand(versionCommand)
}
