package cmd

import (
	"github.com/altopm/alto/utils"
	"github.com/spf13/cobra"
)

var registryCommand = &cobra.Command{
	Use:        "registry",
	Aliases:    []string{},
	SuggestFor: []string{},
	Short:      "Add a new registry",
	Long:       `Add a new registry to the registry list with packages to be installed`,
	Example:    `sudo alto registry add "https://registry.altopkg.com"`,
	ValidArgs: []string{
		"add",
		"remove",
		"list",
	},
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "add" {
			AddRegistry()
		} else if args[0] == "remove" {
			RemoveRegistry()
		} else if args[0] == "list" {
			ListRegistries()
		}
	},
}

func AddRegistry() {

}

func RemoveRegistry() {

}

func ListRegistries() {
	utils.GetRegistryList()
}

func init() {
	rootCmd.AddCommand(registryCommand)
}
