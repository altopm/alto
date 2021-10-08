package cmd

import (
	"fmt"
	"os"

	"github.com/altopm/alto/errors"
	"github.com/altopm/alto/utils"
	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use:     "init",
	Short:   "Initialize a new project",
	Long:    `Initialize a new project to be published to the alto registry, using "sudo alto publish -now"`,
	Example: "sudo alto init altopm",
	Run: func(cmd *cobra.Command, args []string) {
		initPackage(args[0])
	},
}

func initPackage(packageTitle string) {
	initWheel := utils.Loader("%s Initializing...")
	initWheel.Start()
	for i := 0; i < 4; i++ {
		err := os.MkdirAll(fmt.Sprintf("/var/alto/tmp/%s", packageTitle), 0755)
		if err != nil {
			initWheel.Stop()
			errors.Handle(err.Error())
		}
	}
	for i := 0; i < 2; i++ {
		err := os.RemoveAll(fmt.Sprintf("/var/alto/tmp/%s", packageTitle))
		if err != nil {
			initWheel.Stop()
			errors.Handle(err.Error())
		}
	}
	err := os.MkdirAll(".alton", 0755)
	if err != nil {
		initWheel.Stop()
		errors.Handle(err.Error())
	}
	initWheel.Stop()
	utils.MessageSuccess(fmt.Sprintf("%s initialized successfully", packageTitle))
	utils.MessageNeutral("We strongly suggest adding .alton directory to your .gitignore")
}
func init() {
	rootCmd.AddCommand(initCommand)
}
